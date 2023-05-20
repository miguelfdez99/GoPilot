
import { CommandExists } from "../../wailsjs/go/backend/Backend";

export async function checkCommand(command: string) {
    try {
        const commandExists = await CommandExists(command);
        if (!commandExists) {
            alert(
                `System command '${command}' required for this operation is not installed.` +
                ` Please install it and try again.`
            );
        }
    } catch (err) {
        console.error(err);
        alert(
            `Failed to check if system command '${command}' is installed: ${err.message}`
        );
    }
}

export function openDialog(dialog: { showDialog: boolean, dialogTitle: string, dialogMessage: string }, title: string, message: string) {
    dialog.dialogTitle = title;
    dialog.dialogMessage = message;
    dialog.showDialog = true;
    return dialog;
}

export function closeDialog(dialog: { showDialog: boolean, dialogTitle: string, dialogMessage: string }) {
    dialog.dialogTitle = '';
    dialog.dialogMessage = '';
    dialog.showDialog = false;
    return dialog;
}





import { CommandExists, OpenDialogInfo, OpenDialogError, } from "../../wailsjs/go/backend/Backend";

export async function checkCommand(command: string) {
    try {
        const commandExists = await CommandExists(command);
        if (!commandExists) {
            await OpenDialogError(`System command '${command}' required for this operation is not installed. Please install it and try again.`);
        }
    } catch (err) {
        console.error(err);
        await OpenDialogError(`Error checking for system command '${command}'.`);
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




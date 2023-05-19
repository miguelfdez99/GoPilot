
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


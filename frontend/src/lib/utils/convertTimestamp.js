/**
 * @param {number} timestamp
 */
export function formatUnixTimestamp(timestamp) {
    const date = new Date(timestamp); // Directly use the timestamp in milliseconds
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');
    const milliseconds = date.getMilliseconds().toString().padStart(2, '0'); // Add milliseconds
    return `${hours}:${minutes}:${seconds}:${milliseconds}`;
}

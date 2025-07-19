// utils/downloadFile.ts
export async function downloadFile(url: string, fileName: string) {
    const response = await fetch(url);


    console.log(response);
    // Print all headers
    for (let [key, value] of response.headers.entries()) {
        console.log(`${key}: ${value}`);
    }

    if (!response.ok) {
        throw new Error(`Download failed with status ${response.status}`);
    }

    const blob = await response.blob();
    const blobUrl = URL.createObjectURL(blob);

    const link = document.createElement("a");
    link.href = blobUrl;
    link.download = fileName;
    document.body.appendChild(link);
    link.click();
    link.remove();
    URL.revokeObjectURL(blobUrl); // Cleanup
}

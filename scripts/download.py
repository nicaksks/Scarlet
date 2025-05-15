import requests
import os

PATCH_VERSION = "11.2.31"
PLATFORM = "Windows64"
BASE_URL = "http://na-master.projectlumia.com:10800"


def main() -> None:
    request = requests.get(
        f"{BASE_URL}/LIVE/{PATCH_VERSION}/{PLATFORM}/Standalone{PLATFORM}.txt")

    if request.status_code != 200:
        print(f'main: [ERROR] -> {request.reason}')
        return

    getFileName(request.text)


def getFileName(filesName: str) -> None:
    files = filesName.split("\r\n")

    for fileName in files:
        if fileName.endswith('.unity3d'):
            print(f'Starting download file: {fileName}')
            getFileUnity(fileName)


def getFileUnity(fileName: str) -> None:
    request = requests.get(
        f"{BASE_URL}/LIVE/{PATCH_VERSION}/{PLATFORM}/{fileName}")

    if request.status_code != 200:
        print(f'getFileUnity: [ERROR] -> {request.reason}')
        return

    saveFileUnity(fileName, request.content)


def saveFileUnity(fileName: str, content: bytes) -> None:

    if not os.path.exists("./assets"):
        os.mkdir('./assets')

    with open(f'./assets/{fileName}', 'wb') as file:
        try:
            file.write(content)
            print(f'Finishing download file: {fileName}')
            file.close()
        except BlockingIOError as e:
            print(f'saveFileUnity: [ERROR] -> {e}')


if __name__ == "__main__":
    main()

#!/usr/bin/env python
import tempfile
import asarPy
import shutil
import json
import sys
import os


def main() -> None:
    asar_path = sys.argv[1]
    extract_dir = tempfile.mktemp()

    asarPy.extract_asar(asar_path, extract_dir)

    with open(os.path.join(extract_dir, "package.json"), "r", encoding="utf-8") as f:
        package_json = json.load(f)

    package_json["desktopName"] = os.getenv("FLATPAK_ID") + ".desktop"

    with open(os.path.join(extract_dir, "package.json"), "w", encoding="utf-8") as f:
        json.dump(package_json, f)

    asarPy.pack_asar(extract_dir, asar_path)

    shutil.rmtree(extract_dir)


if __name__ == "__main__":
    main()

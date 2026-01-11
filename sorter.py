import os

files = os.listdir()

ignored_files = [".gitignore", "sorter.py", "README.md", "main", "main.go", "go.mod", "go.sum"]

for file in files:
    if file in ignored_files:
        continue

    name, ext = os.path.splitext(file)
    ext = ext[1:]

    if ext == '':
        continue

    if os.path.exists(ext):
        os.rename(file, f"{ext}/{file}")
    else:
        os.makedirs(ext)
        os.rename(file, f"{ext}/{file}")
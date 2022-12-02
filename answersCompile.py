with open("answers.txt", "r") as h:
    data = h.read().split()
    compiled = []
    print(data)
    for i in data:
        i = i.replace("\"", "")
        try:
            compiled.append(str(int(i))+",\n")
            print(f"\nSuccessfully compiled {i}")
        except ValueError:
            print(".", end="")

with open("answers.txt", "w") as h:
    h.writelines(compiled)
file1 = open("file1.txt", "r").readlines()
file2 = open("file2.txt", "r").readlines()

for i in range(len(file1)):
    if file1[i] != file2[i] and i < 1000:
        print(i)
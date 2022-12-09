# Против лома нет приема, если нет другого лома

import time

now = time.localtime(time.time())

print(f"[PYLOM] Encrypting with code {now.tm_hour}")

with open("decrypted.py", "rb") as handler:
    with open("cryptoCntrea.b", "wb") as ready:
        data = handler.read()
        for i in data:
            ready.write((chr(i + now.tm_hour).encode("utf-8")))

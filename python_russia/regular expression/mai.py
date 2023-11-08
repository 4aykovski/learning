import re

text = """asdas 1231 asd 1232, asd... 1231 asd 2222"""

ints = re.findall(r'\d{4}', text)
print(ints)

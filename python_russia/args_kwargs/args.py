a, *b = 1,2,3

def example(a, b, /, c, *, d):
    print(a)
    print(b)
    print(c)
    print(d)

def my_print(*args, **kwargs):
    for arg in args:
        print(str(arg),)
        
        
if __name__ == '__main__':
    my_print(1, 2, 3, 4, sep=':', end = ' ', tr = ' ')    
    #print(1, 2, 3, 4, sep = ':', end ='-')
    
    
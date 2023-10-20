def logger(func):
    """
    Декоратор, который выводит в консоль информацию о том, что
    было передано в функцию.
    """
    def wrapper(*args):
        print(f'{args} were passed to {func.__name__}')
        result = func(*args)
        return result
    
    return wrapper

@logger
def summ(*args: int):
    return sum(args)
    
    
if __name__ == '__main__':
    print(summ(1, 2, 3, 4))

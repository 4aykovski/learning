from contextlib import contextmanager

class Resource:
    def __init__(self):
        self.opened = False
        
    def open(self, *args):
        print(f'Resource was opened with arguments {args}')
        self.opened = True
    
    def close(self):
        print(f'Resource was closed.')
        self.opened = False
        
    def __del__(self):
        if self.opened:
            print('Resource is still opened!')
            
    def action(self):
        print('Do something with resource')
        
@contextmanager
def open_resource(*args):
    resource = None
    try:
        resource = Resource()
        resource.open(1,2,3)
        yield resource
    except Exception:
        raise
    finally:
        if resource:
            resource.close()        
        

if __name__ == '__main__':
    with open_resource(1,2,3) as res:
        res.action()    
        raise ValueError('Some error')
    
import threading
import time


def get_data(data, value):
    for _ in range(value):
        print(f'{threading.current_thread().name} got data {data}')
        time.sleep(1)


if __name__ == '__main__':

    thr_list = []
    for i in range(6):
        thr = threading.Thread(target=get_data, args=(str(time.time()), i,), name=f'thr{i}')
        thr_list.append(thr)
        thr.start()

    print('start')
    for i in thr_list[::-1]:
        i.join()
    print('end')

    # print(f'main thread {threading.main_thread().name}')
    # threading.main_thread().name = 'result'
    # print(f'main thread {threading.main_thread().name}')

#     t1 = threading.Thread(target=get_data, args=(str(time.time()),), name='t1')
#     t1.start()
#
# for i in range(100):
#     print(f"{i} - count")
#     time.sleep(1)
#     if i % 10 == 0:
#         print("active thread", threading.active_count())
#         print("threads - ", threading.enumerate())
#         print('t1 is alive: ', t1.is_alive())
#

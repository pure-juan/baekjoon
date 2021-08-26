def main(number):
    count = 0

    for num in range(1, number + 1):
        if num < 100:
            count += 1
            continue
        temp = list(str(num))
        flag = True
        diff = None
        for index in range(len(temp) - 1):
            if diff == None:
                diff = int(temp[index + 1]) - int(temp[index])
            else:
                if int(temp[index + 1]) - int(temp[index]) != diff:
                    flag = False
                    break

        if flag:
            count += 1

    return count


number = int(input())

print(main(number))

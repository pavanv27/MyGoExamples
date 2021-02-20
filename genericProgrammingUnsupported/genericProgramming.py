
def main():
    # list of string
    s = ["this","is","a","sample","input"]
    print(s)
    reverseListGeneric(s)
    print(s)

    # int list 
    s = [1,2,3,4,5]
    print(s)
    reverseListGeneric(s)
    print(s)

# Generic function which can take any type of list
def reverseListGeneric(s):
    first = 0
    last = len(s) - 1
    while first < last :
        s[first], s[last] = s[last], s[first]
        first+=1
        last-=1

if __name__ == "__main__":
    main()
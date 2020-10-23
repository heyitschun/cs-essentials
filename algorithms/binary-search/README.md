The **binary search** algorithm is an efficient way to search for something in an *ordered* list. The target value is compared to the value in the middle of the list. If they match, then the search ends. If the target is less than the middle element, then the middle value of the first half of the list is compared to the target value. This goes on until there is a match or until the list cannot be halved anymore. The same goes for if the target value is greater than the middle element. 

```
binary_search(arr, target)
    l <- 0
    h <- len(arr)-1
    m <- 0
    
    while l <= h {
        m = (h + l) // 2
        
        if arr[m] < target {
            l = m + 1
        }
        
        else if arr[m] > target {
            h = m - 1
        }
        
        else {
            return m
        }
    }
    
    return -1
```

```
    Depth first 
    parseRoamJSON =
        For each item where item is x:
            if x is the root node:
                write title
                continue

            if x has key children then:
                write parent string to y buffer
                if y is heading:
                    append heading token to y

                write y to document
                parseRoamJSON(x)
```
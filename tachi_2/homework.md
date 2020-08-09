# Question 1
Define an array named "foodz" and set it to random food (Use only lowercase letters).

Ex:
```go
{
    "apple",
    "bologna",
    // etc
}
```

# Question 2
Define an array named "alphabet" and set it to the English alphabet (Use only lowercase letters).

Ex:
```go
// Make sure to use ' and not "
{
    'a'
    'b'
    'c'
    'd'
    'e'
    // etc
}
```

# Question 3
1. Loop through the alphabet array.
1. Println("Examining letter", letter).
1. Inside the alphabet array loop, loop through the food array.
1. If the food starts with the letter, Println(food, "starts with", letter).

Ex:
```go
for _, letter := range alphabet {
    // Println...
    for _, food := range foodz {
        // Logic
            // Println...
    }
}
```

Ex: string logic
```go
if food[0] == []byte(letter)[0]
```

# Question 4
1. Define an array named "numbers" and set it equal to random numbers (Between 0 and 100).
1. Loop through "numbers" and Print all numbers that are less than 50.

# Question 5
Use the rand.Intn() function to set the "numbers" array to random numbers.

Ex:
```go
{
    rand.Intn(100),
    rand.Intn(100),
    rand.Intn(100),
    rand.Intn(100),
    etc
}
```

# Question 6
Define a struct named "Avatar".  Set the leading comment to // Avatar - VRChat avatar.

# Question 7
Add the following properties to Avatar
* Name string
* Author string
* SDKVersion string
* RawData // []byte

# Question 8
Create avatar variables using the "Avatar" struct.

Ex:
```go
loliOnsie := Avatar {
    // Properties
}

cuteCatGirl := Avatar {
    // Properties
}
```

# Question 9
Create an array named "avatars" and set it to the avatars you created.

```go
avatars := {
    avatar1,
    avatar2,
    // etc
}
```

# Question 10
Loop through "avatars" and print them in an informational manner

Ex: 
```go
Println("Avatar name", avatar.Name, "authored by", author) // etc
```
# Slice

An experimental project for do things with very-fresh go's generics - also known as Type Parameter.

## List of methods and usage

- Map

```golang
items := []int{1,2,3,4,5}

result := slice.Map(items, func(i int) int {
  i = i * 2
  return i
})
// the result is []int{2,4,6,8,10}
```

- Reduce

```golang
items := []int{1,2,3,4,5}

result := slice.Reduce(items, func(acc, i int) int{
  return acc + i
}, 0)
// the result is 15
```

- Filter

```golang
// example coming soon
```

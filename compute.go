package main

func Compute(num int64) int64 {

  // Deliberately slow
  if num <= 0 {
    return 0
  } else if num == 1 {
    return 1
  }

  return Compute(num-1) + Compute(num - 2)
}

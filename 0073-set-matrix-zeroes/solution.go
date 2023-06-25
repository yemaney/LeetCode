func setZeroes(matrix [][]int)  {
  x := 1

  for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[0]); j++ {
      if matrix[i][j] == 0 {
        if i == 0 {
          x = 0
          matrix[0][j] = 0
        } else {
        matrix[0][j] = 0
        matrix[i][0] = 0
        }
      }
    }
  }

  // set rows to zero 
  for i := 1; i < len(matrix); i++ {
    if matrix[i][0] == 0 {    
      for j := 1; j < len(matrix[0]); j++ {
        matrix[i][j] = 0
      }
    }
  }

  // set columns to zero 
  for j := 1; j < len(matrix[0]); j++ {
    if matrix[0][j] == 0 {    
      for i := 0; i < len(matrix); i++ {
        matrix[i][j] = 0
      }
    }
  }


  // set first column to all zero's
  if matrix[0][0] == 0 {
    for i := 0; i < len(matrix); i++ {
      matrix[i][0] = 0
    }
  }

  // set first row to all zero's
  if x == 0 {
    for j := 0; j < len(matrix[0]); j++ {
      matrix[0][j] = 0
    }
  }

}

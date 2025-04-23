Program StudentGrades

    Implicit None
    Integer :: NumStudents, I
    Character(Len=30) :: StudentName
    Real :: Grade
    Real, Allocatable :: Grades(:)
    Character(Len=20) :: GradeLevel

    ! Asking The User For The Number Of Students
    Print *, 'Enter The Number Of Students: '
    Read *, NumStudents

    ! Allocating Memory For The Grades Array
    Allocate(Grades(NumStudents))

    ! Reading The Names And Grades Of The Students
    Do I = 1, NumStudents
        Print *, 'Enter The Name Of Student ', I, ': '
        Read *, StudentName
        Print *, 'Enter The Grade For ', StudentName, ': '
        Read *, Grade
        Grades(I) = Grade
    End Do

    ! Calculating And Displaying Grade Levels
    Do I = 1, NumStudents
        If (Grades(I) >= 90) Then
            GradeLevel = 'A'
        Else If (Grades(I) >= 80 .And. Grades(I) < 90) Then
            GradeLevel = 'B'
        Else If (Grades(I) >= 70 .And. Grades(I) < 80) Then
            GradeLevel = 'C'
        Else If (Grades(I) >= 60 .And. Grades(I) < 70) Then
            GradeLevel = 'D'
        Else
            GradeLevel = 'F'
        End If

        ! Displaying Student Information
        Print *, 'Student ', I, ' Has Grade ', Grades(I), ' And Level ', GradeLevel
    End Do

    ! Freeing Allocated Memory
    Deallocate(Grades)

End Program StudentGrades

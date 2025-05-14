Steps :

1- define entities
Entities :
    Task
    Category
    User


2- define properties and behaviors for each Entity

Task:
    Properties:
        -Title
        -DueDate
        -Category
        -IsDone

    Behaviors:
        -Create anew task
        -List User today task
        -List User tasks by Date
        -Change tasks status(done/undone)
        -Edit a task

User:
    Properties:
        -ID
        -Emai
        -Pass

    Behaviors:
        -Register
        -Login

Category:
    Properties:
        -Title
        -Color

    Behaviors:
        -setColor
        -setTitle


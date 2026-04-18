import os
import json
from datetime import datetime
import sys

FILENAME = "tasks.json"

def load_tasks():
    if not os.path.exists(FILENAME):
        return []
    with open(FILENAME, "r") as file:
        try: 
            return json.load(file)
        except json.JSONDecodeError:
            return []
        
def save_tasks(tasks):
    with open(FILENAME, "w") as file:
        json.dump(tasks, file, indent=4)

def add_task(desc):
    tasks = load_tasks()
    task = {
        "id": len(tasks) +1,
        "description": desc,
        "status": "todo",
        "created_at": datetime.now().isoformat(),
        "updated_at": datetime.now().isoformat()
    }
    tasks.append(task)

    save_tasks(tasks)

    print("Add task successfully")

def list_tasks(status=None):
    tasks = load_tasks()

    if status:
        valid_status = ["todo", "in-progress", "done"]
        if status not in valid_status:
            print("Invalid status")
            return
        
        tasks = [task for task in tasks if task["status"]== status]

    if not tasks:
        print("no task found")
        return

    for i in range(len(tasks)):
        print(i, tasks[i])

def delete_task(id):
    tasks = load_tasks()

    new_tasks = [task for task in tasks if task["id"] != id]
    if len(new_tasks) == len(tasks):
        print("Task not found")
        return
    save_tasks(new_tasks)
    print("Deleted Succesfully")

def update_task(id, desc):
    tasks = load_tasks()
    for i in range(len(tasks)):
        if tasks[i]["id"] == id :
            if desc != "":
                tasks[i]["description"] = desc
                tasks[i]["updated_at"] = datetime.now().isoformat()
            save_tasks(tasks)
            print("Successfully updated.")
            return
    print("Task not found")

def change_status(status, id):
    tasks = load_tasks()

    for task in tasks:
        if task["id"] == id:
            task["status"] = status
            task["updated_at"] = datetime.now().isoformat()
            save_tasks(tasks)
            print("Successfully updated")
            return
    print("Task not found")

def main():
    args = sys.argv

    if len(args) < 2:
        print("No command provided")
        return
    
    command = args[1]

    if command == "add":
        desc = args[2]
        add_task(desc)
    elif command == "list":
        if len(args) > 2:
            status = args[2]
            list_tasks(status)
        else:
            list_tasks()
    elif command == "update":
        id = int(args[2])
        desc = args[3]
        update_task(id, desc)
    elif command == "delete":
        id = int(args[2])
        delete_task(id)
    elif command == "mark-done":
        id = int(args[2])
        change_status("done", id)
    elif command == "mark-in-progress":
        id = int(args[2])
        change_status("in-progress", id)
    else:
        print("Unknown command")

if __name__ == "__main__":
    main()
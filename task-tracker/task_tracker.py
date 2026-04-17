import os
import json

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
        "status": "todo"
    }
    tasks.append(task)

    save_tasks(tasks)

    print("Add task successfully")

def list_tasks():
    tasks = load_tasks()

    for i in range(len(tasks)):
        print(i, tasks[i])

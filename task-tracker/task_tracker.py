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
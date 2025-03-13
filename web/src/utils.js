
export const baseUrl = "http://localhost:8080/v1/tasks"

export const createTask = async (newTask) => {
    const res = await fetch(baseUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newTask)
    })

  }

  export const deleteTask = async (id) => {
    const res = await fetch(`${baseUrl}/${id}`, {
      method: "DELETE",
    })
  }

  export const updateTask = async (task) => {
    const res = await fetch(`${baseUrl}/${task.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(task)
    })
  }
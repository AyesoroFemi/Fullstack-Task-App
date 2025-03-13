import { useEffect, useState } from 'react'
import './App.css'
import { baseUrl, createTask, deleteTask, updateTask } from './utils'
import Delete from "./assets/delete.svg"


function App() {
  const [tasks, setTasks] = useState([])
  const [title, setTitle] = useState("")

  // function to get all the tasks
  const fetchTasks = async () => {
    const response = await fetch(baseUrl)
    const data = await response.json()
    setTasks(data.data)
    console.log("getting data", data)
  }

  useEffect(() => {
    fetchTasks()
  }, [])


  // function to delete a task
  const handleDelete = async (taskId) => {
    const confirm = window.confirm("Are you sure you want to delete this task")
    if (!confirm) return
    await deleteTask(taskId)
    fetchTasks()
  }

  
  // function to toggle a task
  const toggleTaskCompletion = async (task) => {
    const updatedTask = {
      ...task,
      completed: task.completed ? 0 : 1
    }
    await updateTask(updatedTask)
    fetchTasks()
  }



// Handle submit function for form submission
  const handleSubmit = async (e) => {
    e.preventDefault()

    if (!title.trim()) {
      alert("Title cannot be empty!")
      return
    }
    const newTask = {
      title: title.trim(),
      completed: 0
    }

    console.log(newTask)
    await createTask(newTask)
    fetchTasks()
    setTitle("")
  }

  return (
    <>
      <div className='container'>
        <h3>Task App</h3>
        <form className='form' onSubmit={handleSubmit}>
          <div className='card__form'>
            <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} />
            <button className='submit' type='submit'>Submit</button>
          </div>
        </form>
        {tasks?.length > 0 ? <div className='task__card'>
          {tasks.map((task) => {
            return (
              <div className='card' key={task.id}>
                <input type="checkbox" checked={task.completed} onChange={() => toggleTaskCompletion(task)} />
                <span className={`card__title ${task.completed ? "task__title" : ""}`} >{task.title}</span>
                <button onClick={() => handleDelete(task.id)} className='del'>
                  <img src={Delete} alt="delete icon" />
                </button>
              </div>
            )
          })}
        </div> : <h3>No task added yet</h3>}
      </div>
    </>
  )
}

export default App

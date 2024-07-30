<script>
  import axios from 'axios';
  export default {
    name: 'Home',
    data() {
      return {
        tasks: []
      }
    },
    methods: {
      getTasks() {
        axios.get('http://localhost:3000/task')
        .then(response => {
            this.tasks = response.data
        }).catch(err => {
            console.log("Couldn't get tasks: "+err)
        })
      }
    },
    created() {
        this.getTasks()
    }
  }
</script>

<template>
<div id="container">
    <div>
        <div class="header">
            <h1>TO-DO List</h1>
            <button id="add-btn">Add Task</button>
        </div>
        <div class="card" v-for="(task,index) in tasks">
            <div class="header">
                {{ index }}
                <h4>{{ task.title }}</h4>
                <button
                    class="material-icons"
                    id="edit-btn">edit</button
                >
            </div>
            <p>{{task.status}}</p>
            <p>{{task.date_limit}}</p>
            <div id="bottom">
                <button id="delete-btn" onclick="deleteTask(task.id)">Delete</button>
            </div>
        </div>
    </div>

</div>
</template>

<style>
  #container {
      display: flex;
      justify-content: center;
      align-items: center;
  }
  #add-btn {
      cursor: pointer;
      background-color: rgb(130, 197, 28);
      border: 3px solid rgb(9, 148, 9);
      width: 110px;
      height: 36px;
      margin-left: 16px;
      font-size: 18px;
      border-radius: 4px;
      font-weight: 600;
  }
  .card {
      width: 300px;
      height: 200px;
      padding-left: 16px;
      padding-right: 16px;
      background-color: whitesmoke;
      color: rgb(22, 22, 22);
      border-radius: 4px;
      margin-bottom: 20px;
  }
  .card:hover {
      box-shadow: rgba(0, 0, 0, 0.479) 0px 8px 20px;
      transition: 0.3s;
  }
  .header {
      display: flex;
      align-items: center;
      justify-content: space-between;
  }
  #edit-btn {
      cursor: pointer;
      border: 1px solid transparent;
      border-radius: 4px;
  }
  #edit-btn:hover {
      border-color: rgb(22, 22, 22);
  }
  #delete-btn {
      font-family: "Trebuchet MS";
      font-weight: 600;
      font-size: 14px;
      cursor: pointer;
      background-color: rgb(201, 59, 59);
      border: 2px solid rgb(133, 37, 37);
      border-radius: 4px;
      width: 72px;
      height: 28px;
  }
  #bottom {
      display: flex;
      justify-content: end;
  }
</style>

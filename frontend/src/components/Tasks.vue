<script>
  import axios from 'axios';
  export default {
    name: 'Home',
    data() {
      return {
        tasks: [],
        editDialog: false,
        addDialog: false,
        title: '',
        status: '',
        selectedId: '',
        snackbar: false,
        snkMsg: ''
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
      },
      addTask() {
        axios.post('http://localhost:3000/task', {
            task: this.title,
            date_limit: '2024-08-12T11:30:00Z'
        }).then(() => {
            this.snkMsg = 'Task added successfully'
            this.snackbar = true
        }).catch((err) => {
            this.snkMsg = "Couldn't add task: "+err
        })
      },
      editTask() {
        axios.patch(`http://localhost:3000/task/${this.selectedId}`, {
            task: this.title,
            status: this.status
        }).then(() => {
            this.snkMsg = 'Task edited successfully'
            this.snackbar = true
        }).catch(err  => {
            this.snackbar = true
            this.snkMsg = 'Edit task error: '+err
        })
      },
      deleteTask(id) {
        axios.delete(`http://localhost:3000/task/${id}`)
        .then(() => {
            this.snkMsg = 'Task deleted successfully'
            this.snackbar = true
        }).catch((err) => {
            this.snkMsg = "Task couldn't be deleted"
            this.snackbar = true
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
            <button id="add-btn" v-on:click="addDialog=true">Add Task</button>
        </div>
        <div class="card" v-for="(task,index) in tasks">
            <div class="header">
                {{ index }}
                <h4>{{ task.title }}</h4>
                <button
                    class="material-icons"
                    id="edit-btn" v-on:click="editDialog=true, selectedId= task.id">edit</button
                >
            </div>
            <p>{{task.status}}</p>
            <p>{{task.date_limit}}</p>
            <div id="bottom">
                <button id="delete-btn" v-on:click="deleteTask(task.id)">Delete</button>
            </div>
        </div>
    </div>
</div>

<!-- Dialog -->
<div class="pa-4 text-center">
<v-dialog v-model="editDialog" max-width="400">

    <v-card title="Edit Task">
    <v-card-text>
        <v-row dense>
            <v-col cols="12" md="10" sm="10">
                <v-text-field
                label="Title"
                v-model="title"
                ></v-text-field>
            </v-col>
        
            <v-col cols="12" md="10" sm="10">
                <v-select
                :items="['Not started', 'In process', 'Completed']"
                label="Status"
                v-model="status"
                auto-select-first
                ></v-select>
            </v-col>
        </v-row>
    </v-card-text>
    <v-divider></v-divider>

    <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn
        text="Close"
        variant="plain"
        @click="editDialog = false"
        ></v-btn>

        <v-btn
        color="primary"
        text="Save"
        variant="tonal"
        @click="editDialog = false, editTask()"
        ></v-btn>
    </v-card-actions>
    </v-card>
</v-dialog>
</div>

<div class="pa-4 text-center">
<v-dialog v-model="addDialog" max-width="400">

    <v-card title="Add Task">
    <v-card-text>
        <v-row dense>
            <v-col cols="12" md="10" sm="10">
                <v-text-field
                label="Title"
                v-model="title"
                ></v-text-field>
            </v-col>
        </v-row>
    </v-card-text>
    <v-divider></v-divider>

    <v-card-actions>
        <v-spacer></v-spacer>

        <v-btn
        text="Close"
        variant="plain"
        @click="addDialog = false"
        ></v-btn>

        <v-btn
        color="primary"
        text="Save"
        variant="tonal"
        @click="addDialog = false, addTask()"
        ></v-btn>
    </v-card-actions>
    </v-card>
</v-dialog>
</div>
<v-snackbar
      v-model="snackbar"
    >
      {{ snkMsg }}

      <template v-slot:actions>
        <v-btn
          color="pink"
          variant="text"
          @click="snackbar = false"
        >Close</v-btn>
      </template>
    </v-snackbar>
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
      padding-top: 16px;
      background-color: whitesmoke;
      color: rgb(22, 22, 22);
      border-radius: 4px;
      margin-top: 20px;
  }
  .card:hover {
      box-shadow: rgba(0, 0, 0, 0.479) 0px 8px 20px;
      transition: 0.3s;
  }
  .header {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 12px;
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

<template>
  <v-card class="ma-5 pa-5 grey lighten-4">
    <div v-if="message" class="alert alert-success">
      {{ this.message }}</div>
    <div class="container">
      <table class="table">
        <thead>
          <tr>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Email</th>
            <th>Update</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" v-bind:key="user.id">
            <td>{{ user.firstName }}</td>
            <td>{{ user.lastName }}</td>
            <td>{{ user.email }}</td>
            <td>
              <v-btn class="btn btn-warning" @click="updateUser(user.id)">
                Update
              </v-btn>
            </td>
            <td>
              <v-btn color="red" @click="deleteUser(user.id)">
                Delete
              </v-btn>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="row">
        <v-btn color="green" @click="addUser()">Add</v-btn>
      </div>
    </div>
  </v-card>
</template>

  
<script>
import UserDataService from "../service/UserDataService";

export default {
  name: "Users",
  data() {
    return {
      users: [],
      message: "",
    };
  },
  methods: {
    refreshUsers() {
      UserDataService.retrieveAllUsers().then((res) => {
        this.users = res.data;
      });
    },
    addUser() {
      this.$router.push(`/user/-1`);
    },
    updateUser(id) {
      this.$router.push(`/user/${id}`);
    },
    deleteUser(id) {
      UserDataService.deleteUser(id).then(() => {
        this.refreshUsers();
      });
    },
  },
  created() {
    this.refreshUsers();
  },
};
</script>
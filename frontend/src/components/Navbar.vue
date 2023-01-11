<template>

    <nav>

        <v-toolbar flat app>
            <v-app-bar-nav-icon @click="drawer = !drawer" class="grey--text"></v-app-bar-nav-icon>
            <v-toolbar-title class="text-uppercase grey--text">
                <span class="font-weight-light">Ishmam </span>
                <span>Abir</span>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn flat color="grey">
                <span>Test Button</span>
                <v-icon right>mdi-account</v-icon>
            </v-btn>
        </v-toolbar>

        <v-navigation-drawer app v-model="drawer" class="indigo">
            <v-list dense nav>
                <v-list-item link :to="{path: '/users'}">
                    <v-list-item-icon>
                        <v-icon>mdi-view-dashboard</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                        <v-list-item-title>All Users</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-list-item link :to="{ path: '/user/-1' }">
                    <v-list-item-icon>
                        <v-icon>mdi-account-multiple</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                        <v-list-item-title>Add New User</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
            <br><br><br><br>
            <v-btn @click="(drawer = !drawer)" color="green">CLose</v-btn>
        </v-navigation-drawer>

    </nav>

</template>
  
<script>
export default {
    name: 'Navbar',

    data() {
        return {
            drawer: false
        }
    }
}
</script>
  
<style>

</style>
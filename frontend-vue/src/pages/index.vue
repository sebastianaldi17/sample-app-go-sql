<template>
    <v-navigation-drawer v-model="drawer" temporary>
        <v-list-item :title="`Hello, ${username}`"></v-list-item>
        <v-divider></v-divider>
        <v-list-item prepend-icon="mdi-note-plus" @click="redirectNewTodo">
            New ToDo
        </v-list-item>
        <v-list-item prepend-icon="mdi-account" @click="logout">
            Logout
        </v-list-item>
    </v-navigation-drawer>

    <v-app-bar>
        <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-app-bar-title>ToDo App</v-app-bar-title>
    </v-app-bar>

    <v-main class="pt-2">
        <v-container>
            <v-row v-if="!isFetching">
                <v-col cols="6" sm="4" md="3" lg="2" v-for="todo in todos">
                    <v-card :title="todo.title" :text="todo.content"
                        :subtitle="`${todo.completed ? 'done' : 'not done'}`"
                        @click="redirectEditTodo(todo.id)">
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </v-main>
</template>

<script lang="ts" setup>
import { Todo } from '@/assets/Itodo';
import Utils from '@/assets/utils';
import router from '@/router';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const drawer = defineModel<boolean>("drawer")

let todos: Todo[] = []

let username: string

let isFetching = ref(true)

onMounted(() => {
    let token = sessionStorage.getItem("jwt")
    const router = useRouter()
    if (token === null) {
        router.push("/login")
        return
    }

    axios.get(`${import.meta.env.VITE_API_URL}/user`, {
        headers: { Authorization: `Bearer ${token}` }
    })
        .catch(() => {
            sessionStorage.removeItem("jwt")
            router.push("/login")
            return
        })

    let jwtPayload = Utils.parseJwt(token)
    username = jwtPayload.username

    axios.get(`${import.meta.env.VITE_API_URL}/user/todo`, {
        headers: { Authorization: `Bearer ${token}` }
    }).then((response: AxiosResponse<Todo[]>) => {
        console.log(response.data)
        for (let i = 0; i < response.data.length; i += 1) {
            todos.push({
                completed: response.data[i].completed,
                content: response.data[i].content,
                title: response.data[i].title,
                last_update: response.data[i].last_update,
                created_at: response.data[i].created_at,
                id: response.data[i].id,
            } as Todo)
        }

        isFetching.value = false
        return
    })
        .catch((error: AxiosError) => {
            console.error(error)
            alert("An error occured, please try again later or relogin")
        })
})

function logout() {
    sessionStorage.removeItem("jwt")
    window.location.reload()
}

function redirectNewTodo() {
    router.push("/todo/new-todo")
}

function redirectEditTodo(todoID: number) {
    router.push(`/todo/${todoID}`)
}

</script>

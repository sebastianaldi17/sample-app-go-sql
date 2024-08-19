<template>
    <v-navigation-drawer v-model="drawer" temporary>
        <v-list-item :title="`Hello, ${username}`"></v-list-item>
        <v-divider></v-divider>
        <v-list-item prepend-icon="mdi-home" @click="redirectHome">
            Home
        </v-list-item>
        <v-list-item prepend-icon="mdi-account" @click="logout">
            Logout
        </v-list-item>
    </v-navigation-drawer>

    <v-app-bar>
        <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-app-bar-title>Edit ToDo</v-app-bar-title>
    </v-app-bar>

    <v-main class="pt-2">
        <v-container>
            <v-form v-model="formModel">
                <v-row>
                    <v-col cols="12">
                        <v-text-field label="Title" prepend-icon="mdi-format-title" hide-details="auto"
                            v-model="titleModel" :rules="[rules.required]">
                        </v-text-field>
                    </v-col>
                    <v-col cols="12">
                        <v-text-field label="Content" prepend-icon="mdi-text" hide-details="auto" v-model="contentModel"
                            :rules="[rules.required]">
                        </v-text-field>
                    </v-col>
                    <v-col cols="12">
                        <v-checkbox label="Completed?" v-model="completedModel"></v-checkbox>
                    </v-col>
                    <v-col cols="12">
                        <v-btn :disabled="!formModel" @click="editToDo" color="primary">Edit ToDo</v-btn>
                        <v-btn @click="deleteToDo" color="red-accent-4" class="ml-4">Delete ToDo</v-btn>
                    </v-col>
                </v-row>
            </v-form>
        </v-container>
    </v-main>
</template>

<script lang="ts" setup>
import { Todo } from '@/assets/Itodo';
import Utils from '@/assets/utils';
import router from '@/router';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const drawer = defineModel<boolean>("drawer")
const titleModel = defineModel<string>("titleModel")
const contentModel = defineModel<string>("contentModel")
const completedModel = defineModel<boolean>("completedModel")
const formModel = defineModel<boolean>("formModel")

const rules = {
    required: function (value: string) {
        return !!value || "This field is required"
    },
}

const route = useRoute()

let todoID = Number(route.params.id)

let username: string

onMounted(() => {
    const router = useRouter()
    if (isNaN(todoID)) {
        router.push("/")
        return
    }

    let token = sessionStorage.getItem("jwt")
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

    axios.get(`${import.meta.env.VITE_API_URL}/todo/${todoID}`, {
        headers: { Authorization: `Bearer ${token}` }
    })
        .then((response: AxiosResponse<Todo>) => {
            titleModel.value = response.data.title
            contentModel.value = response.data.content
            completedModel.value = response.data.completed
        })
        .catch((error: AxiosError) => {
            alert("An error occured, please try again later")
            console.error(error)
        })
})

function logout() {
    sessionStorage.removeItem("jwt")
    window.location.reload()
}

function redirectHome() {
    router.push("/")
}

function editToDo() {

    let token = sessionStorage.getItem("jwt")
    let title = titleModel.value
    let content = contentModel.value
    let completed = completedModel.value
    if (title === undefined || content === undefined || token === null || completed == null) {
        return
    }

    axios.put(`${import.meta.env.VITE_API_URL}/todo/${todoID}`, {
        title: title,
        content: content,
        completed: completed,
    }, {
        headers: { Authorization: `Bearer ${token}` }
    }).then(() => {
        router.push("/")
    }).catch((error: AxiosError) => {
        console.error(error)
        alert("An error occured, please try again later")
    })

}
function deleteToDo() {

    let token = sessionStorage.getItem("jwt")
    if (token === null) {
        return
    }

    axios.delete(`${import.meta.env.VITE_API_URL}/todo/${todoID}`, {
        headers: { Authorization: `Bearer ${token}` }
    }).then(() => {
        router.push("/")
    }).catch((error: AxiosError) => {
        console.error(error)
        alert("An error occured, please try again later")
    })

}
</script>
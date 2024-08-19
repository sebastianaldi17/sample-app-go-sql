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
        <v-app-bar-title>New ToDo</v-app-bar-title>
    </v-app-bar>

    <v-main class="pt-2">
        <v-container>
            <v-form v-model="formModel">
                <v-row>
                    <v-col cols="12">
                        <v-text-field label="Title" prepend-icon="mdi-format-title" hide-details="auto" v-model="titleModel" :rules="[rules.required]">
                        </v-text-field>
                    </v-col>
                    <v-col cols="12">
                        <v-text-field label="Content" prepend-icon="mdi-text" hide-details="auto" v-model="contentModel" :rules="[rules.required]">
                        </v-text-field>
                    </v-col>
                    <v-col>
                        <v-btn :disabled="!formModel" @click="createToDo">Create ToDo</v-btn>
                    </v-col>
                </v-row>
            </v-form>
        </v-container>
    </v-main>
</template>

<script lang="ts" setup>
import Utils from '@/assets/utils';
import router from '@/router';
import axios, { AxiosError } from 'axios';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

const drawer = defineModel<boolean>("drawer")
const titleModel = defineModel<string>("titleModel")
const contentModel = defineModel<string>("contentModel")
const formModel = defineModel<boolean>("formModel")

const rules = {
    required: function (value: string) {
        return !!value || "This field is required"
    },
}

let username: string

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
})

function logout() {
    sessionStorage.removeItem("jwt")
    window.location.reload()
}

function redirectHome() {
    router.push("/")
}

function createToDo() {
    let token = sessionStorage.getItem("jwt")
    let title = titleModel.value
    let content = contentModel.value
    if (title === undefined || content === undefined || token === null) {
        return
    }

    axios.post(`${import.meta.env.VITE_API_URL}/todo`, {
        title: title,
        content: content,
    }, {
        headers: { Authorization: `Bearer ${token}` }
    }).then(() => {
        router.push("/")
    })
    .catch((error: AxiosError) => {
        console.error(error)
        alert("An error occured, please try again later")
    })
}

</script>
<template>
    <v-container fluid class="fill-height">
        <v-row class="d-flex justify-center align-center">
            <v-col cols="12" sm="8" md="4">
                <v-card>
                    <v-toolbar dark color="primary">
                        <v-toolbar-title>Login</v-toolbar-title>
                    </v-toolbar>
                    <v-form v-model="formModel">
                        <v-card-text>
                            <v-text-field prepend-icon="mdi-account" label="Username" type="text" hide-details="auto"
                                v-model="usernameModel"></v-text-field>
                        </v-card-text>
                        <v-card-text>
                            <v-text-field prepend-icon="mdi-lock" label="Password" type="password" hide-details="auto"
                                v-model="passwordModel"></v-text-field>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="primary" href="/register">Register</v-btn>
                            <v-btn color="secondary" @click="loginUser">Login</v-btn>
                        </v-card-actions>
                    </v-form>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts" setup>
import router from '@/router';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { onBeforeMount } from 'vue';

onBeforeMount(() => {
    let token = sessionStorage.getItem("jwt")
    if (token !== null) {
        axios.get(`${import.meta.env.VITE_API_URL}/user`, {
            headers: { Authorization: `Bearer ${token}` }
        })
            .then(() => {
                router.push("/")
            })
            .catch((error) => {
                console.error(error)
                sessionStorage.removeItem("jwt")
                return
            })
    }
})

const usernameModel = defineModel<string>("usernameModel")
const passwordModel = defineModel<string>("passwordModel")
const formModel = defineModel<boolean>("formModel")

function loginUser() {
    let username = usernameModel.value
    let password = passwordModel.value
    if (username === undefined || password === undefined) {
        return
    }

    axios.post(`${import.meta.env.VITE_API_URL}/user/login`, {
        username: username,
        password: password
    })
        .then((response: AxiosResponse) => {
            sessionStorage.setItem("jwt", response.data.token)
            router.push("/")
        })
        .catch((error: AxiosError) => {
            console.error(error)
            if (error.response?.status === 401) {
                alert("Invalid username or password")
            } else {
                alert("An error occured, please try again later")
            }
        })
}

</script>
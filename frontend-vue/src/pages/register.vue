<template>
    <v-container fluid class="fill-height">
        <v-row class="d-flex justify-center align-center">
            <v-col cols="12" sm="8" md="4">
                <v-card>
                    <v-toolbar dark color="primary">
                        <v-toolbar-title>Register</v-toolbar-title>
                    </v-toolbar>
                    <v-form v-model="formModel">
                        <v-card-text>
                            <v-text-field prepend-icon="mdi-account" label="Username" type="text" hide-details="auto"
                                v-model="usernameModel" :rules="[rules.required, rules.usernameLength]"></v-text-field>
                        </v-card-text>
                        <v-card-text>
                            <v-text-field prepend-icon="mdi-lock" label="Password" type="password"
                                v-model="passwordModel" hide-details="auto"
                                :rules="[rules.required, rules.passwordLength]"></v-text-field>
                        </v-card-text>
                        <v-card-text>
                            <v-text-field prepend-icon="mdi-lock" label="Confirm Password" type="password"
                                v-model="confirmPasswordModel" hide-details="auto"
                                :rules="[rules.required, rules.samePassword]"></v-text-field>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="secondary" @click="goToLoginPage">Login</v-btn>
                            <v-btn color="primary" @click="createAccount" :disabled="!formModel">Create Account</v-btn>
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
const confirmPasswordModel = defineModel<string>("confirmPasswordModel")
const formModel = defineModel<boolean>("formModel")

const rules = {
    required: function (value: string) {
        return !!value || "This field is required"
    },

    samePassword: function (value: string) {
        return value === passwordModel.value || "Password must be the same"
    },

    usernameLength: function (value: string) {
        return (value.length >= 8 && value.length <= 16) || "Username must be between 8-16 characters"
    },

    passwordLength: function (value: string) {
        return (value.length >= 8 && value.length <= 16) || "Password must be between 8-16 characters"
    }
}

function createAccount() {
    let username = usernameModel.value
    let password = passwordModel.value
    let confirmPassword = confirmPasswordModel.value
    if (username === undefined || password === undefined || confirmPassword === undefined) {
        return
    }

    axios.post(`${import.meta.env.VITE_API_URL}/user`, {
        username: username,
        password: password
    })
        .then((response: AxiosResponse) => {
            if (response.status !== 200) {
                alert("An error occured, plesae try again later")
                return
            }
            router.push("/login")
        })
        .catch((error: AxiosError) => {
            console.error(error)
            alert("An error occured, please try again later")
        })
}

function goToLoginPage() {
    router.push("/login")
}
</script>
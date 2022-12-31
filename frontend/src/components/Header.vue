<template>
    <div id="header" class="row align-center">
        <div class="flex xs3">
            <div class="item"></div>
        </div>
        <div class="flex xs6">
            <div class="item">File Integrity Verifier</div>
        </div>
        <div class="flex xs1">
            <div class="item"></div>
        </div>
        <div class="flex xs2">
            <div class="item"><va-button size="small" @click="addfiles">Add File</va-button></div>
        </div>
    </div>

</template>

<style scoped>
#header {
    padding-top: 2rem;
    padding-bottom: 2rem;
}
</style>

<script lang="ts" setup>
import { AddFile } from "../../wailsjs/go/main/App.js";

function addfiles() {
    AddFile()
}
</script>
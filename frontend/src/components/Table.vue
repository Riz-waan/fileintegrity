<template>
  <div class="va-table-responsive">
    <table class="va-table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Path</th>
          <th>Last Successful Update</th>
          <th>Last Update</th>
          <th>Status</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in state.users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.path }}</td>
          <td>{{ (new Date(Date.parse(user.successfulUpdateAt))).toLocaleString() }}</td>
          <td>{{ (new Date(Date.parse(user.updatedAt))).toLocaleString() }}</td>
          <td>
            <va-badge :color="user.successfulUpdateAt === user.updatedAt ? 'success' : 'danger'"
              :text="user.successfulUpdateAt === user.updatedAt ? 'intact' : 'compromised'" />
          </td>
          <td> <va-button size="small" @click="deletefile(user.id)">Remove</va-button></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { GetFiles, RemoveFile } from "../../wailsjs/go/main/App.js";
import { onMounted } from 'vue'
import { EventsOn } from "../../wailsjs/runtime"

const users = [{
  id: 1,
  path: 'Loading',
  successfulUpdateAt: '',
  updatedAt: '',
}]

const state = reactive({ users })

function getfiles() {
  GetFiles().then((result) => (state.users = result));
}
function deletefile(id: number) {
  RemoveFile(id)
}

onMounted(() => {
  getfiles()
})

EventsOn("refetch", () => getfiles())


</script>

<style lang="scss" scoped>
.va-table-responsive {
  overflow: auto;
}

.va-table {
  width: 100%;
}

.va-table td {
  /* <http://www.w3.org/wiki/CSS/Properties/text-align>
     * left, right, center, justify, inherit
     */
  text-align: left;
  /* <http://www.w3.org/wiki/CSS/Properties/vertical-align>
     * baseline, sub, super, top, text-top, middle,
     * bottom, text-bottom, length, or a value in percentage
     */
  vertical-align: middle;
}
</style>

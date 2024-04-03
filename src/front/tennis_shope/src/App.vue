<template>
    <h1>{{info}}</h1>
    <button type="button" @click="userData()">Отправить</button>
</template>

<script>
export default {
    data() {
        return {
            info: 'Привет'
        }
    },

    methods : {
        userData() {
            this.info = 'Как дела'
        }
    }


}

</script>

<style scoped>

</style>
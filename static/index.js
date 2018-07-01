var defaultDate = moment().add(1, 'days').format("YYYY-MM-DD")

var app = new Vue({
    el: '#app',
    data: {
        reservations: null,
        date: defaultDate,
        name: "",
        min: defaultDate 
    },mounted () {
        axios
        .get('http://localhost:8080/api/reservations')
        .then(response => {
            this.reservations = response.data
        })
    }, methods:{
        create: function (event) {
            var fd = new FormData()
            fd.set("name", this.name)
            fd.set("date", new Date(this.date).toISOString())
            axios({
                method: 'put',
                url: 'http://localhost:8080/api/reservations',
                data: fd,
                config: { headers: {'Content-Type': 'multipart/form-data' }}
            })
            .then(response => {
                console.log(response)
            })
            event.preventDefault()
        }
    }
})

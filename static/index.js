var defaultDate = moment().add(1, 'days').format("YYYY-MM-DD")

function update(ctx) {
    axios
        .get('http://localhost:8080/api/reservations')
        .then(response => {
            ctx.validReservations = response.data.valid_reservations
        })
}

var app = new Vue({
    el: '#app',
    data: {
        validReservations: null,
        expiredReservations: null,
        date: defaultDate,
        name: "",
        min: defaultDate
    },mounted () {
        update(this)
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
                update(this)
            })
            event.preventDefault()
        }
    }
})

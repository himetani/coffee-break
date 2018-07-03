var defaultDate = moment().add(1, 'days').format("YYYY-MM-DD")

function update(ctx) {
    axios
        .get('http://localhost:8080/api/reservations')
        .then(response => {
            ctx.validReservations = response.data.valid_reservations
        })
}

function hideNotificationFn(ctx) {
    return function() {
        ctx.success = false
        ctx.fail = false
    }
}


var app = new Vue({
    el: '#app',
    data: {
        validReservations: null,
        expiredReservations: null,
        date: defaultDate,
        name: "",
        min: defaultDate,
        success: false,
        fail: false,
        errorMsg: ""
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
                this.success = true
                this.date = defaultDate
                var hideNodification = hideNotificationFn(this)
                setTimeout(hideNodification, 1000)
            })
            event.preventDefault()
        }
    }
})

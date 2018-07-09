var defaultDate = moment().add(1, 'days').format("YYYY-MM-DD")

function update(ctx) {
    axios
        .get('http://localhost:8080/api/reservations')
        .then(response => {
            ctx.validReservations = response.data.valid_reservations
        })
}

function toSubmitFn(ctx) {
    return function() {
        ctx.buttonType = 'submit'
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
        buttonType: "submit"
    },mounted () {
        update(this)
    }, methods:{
        create: function (event) {
            if(this.buttonType !== "submit") {
                event.preventDefault()
                return
            }
            this.buttonType = "processing"
            var fd = new FormData()
            fd.set("name", this.name)
            fd.set("date", new Date(this.date).toISOString())
            axios({
                method: 'post',
                url: 'http://localhost:8080/api/reservations',
                data: fd,
                config: { headers: {'Content-Type': 'multipart/form-data' }}
            })
            .then(response => {
                console.log(response)
                update(this)
                this.date = defaultDate
                this.buttonType = 'success'
                this.name = ""
                var toSubmit = toSubmitFn(this)
                setTimeout(toSubmit, 5000)
            }).catch(error => {
                this.date = defaultDate
                this.buttonType = 'fail'
                this.name = ""
                var toSubmit = toSubmitFn(this)
                setTimeout(toSubmit, 5000)
            });
            event.preventDefault()
        }
    }
})

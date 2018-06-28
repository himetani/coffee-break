var defaultDate = moment().add(1, 'days').format("YYYY-MM-DD");

var app = new Vue({
    el: '#app',
    data: {
        reservations: null,
        date: defaultDate,
        name: ""
    },mounted () {
        axios
        .get('http://localhost:8080/api/reservations')
        .then(response => {
            this.reservations = response.data
        })
    }, methods:{
        addReservation: function (event) {
            console.log(this.date, this.name)
        }
    }
})

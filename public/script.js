Vue.createApp({
  data() {
    return {
      description:
        'This is a Golang REST API built with Go Fiber and using a statically hosted Vue.js front-end',
      jsonData: null,
    };
  },
  methods: {
    getJson() {
      axios.get('/json').then((response) => {
        this.jsonData = response.data;
      });
    },
  },
}).mount('#app');

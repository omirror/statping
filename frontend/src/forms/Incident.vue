<template>
<div>

    <div v-for="(incident, i) in incidents" class="card contain-card text-black-50 bg-white mb-4">
        <div class="card-header">Incident: {{incident.title}}
            <button @click="deleteIncident(incident)" class="btn btn-sm btn-danger float-right">
                <font-awesome-icon icon="times" />  Delete
            </button>
        </div>
                <div class="card-body bg-light pt-3">
                    <div v-for="(update, i) in incident.updates" class="alert alert-light" role="alert">
                        <span class="badge badge-pill badge-info text-uppercase">{{update.type}}</span>
                        <span class="float-right font-2">{{ago(update.created_at)}} ago</span>
                        <span class="d-block mt-2">{{update.message}}
                        <button @click="delete_update(update)" type="button" class="close" data-dismiss="alert" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                        </span>
                    </div>

                    <FormIncidentUpdates :incident="incident"/>

                    <span class="font-2 mt-3">Created: {{niceDate(incident.created_at)}} | Last Update: {{niceDate(incident.updated_at)}}</span>
                </div>
    </div>


    <div class="card contain-card text-black-50 bg-white mb-5">
        <div class="card-header">Create Incident for {{service.name}}</div>
        <div class="card-body">
            <form @submit.prevent="createIncident">
                <div class="form-group row">
                    <label class="col-sm-4 col-form-label">Title</label>
                    <div class="col-sm-8">
                        <input v-model="incident.title" type="text" name="title" class="form-control" id="title" placeholder="Incident Title" required>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-sm-4 col-form-label">Description</label>
                    <div class="col-sm-8">
                        <textarea v-model="incident.description" rows="5" name="description" class="form-control" id="description" required></textarea>
                    </div>
                </div>

                <div class="form-group row">
                    <div class="col-sm-12">
                        <button @click.prevent="createIncident"
                                :disabled="!incident.title || !incident.description"
                                type="submit" class="btn btn-block btn-primary">
                            Create Incident
                        </button>
                    </div>
                </div>
                <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
            </form>
            </div>
    </div>
</div>
</template>

<script>
  import Api from "../API";
  import flatPickr from 'vue-flatpickr-component';
  import 'flatpickr/dist/flatpickr.css';
  import FormIncidentUpdates from './IncidentUpdates';

  export default {
  name: 'FormIncident',
  components: {
      FormIncidentUpdates
  },
  props: {
    service: {
      type: Object
    }
  },
  data () {
    return {
      incident: {
        title: "",
        description: "",
        service: this.service.id,
      },
        incidents: [],
    }
  },
      async mounted () {
          await this.loadIncidents()
      },
      methods: {
        async delete_update(update) {
          await Api.incident_update_delete(update)
          this.incidents = await Api.incidents_service(this.service)
        },
        async loadIncidents() {
          this.incidents = await Api.incidents_service(this.service)
        },
          async createIncident() {
            await Api.incident_create(this.service, this.incident)
            await this.loadIncidents()
            this.incident = {
                title: "",
                description: "",
                service: this.service.id,
            }
          },
      async deleteIncident(incident) {
          let c = confirm(`Are you sure you want to delete '${incident.title}'?`)
          if (c) {
              await Api.incident_delete(incident)
            await this.loadIncidents()
          }
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>

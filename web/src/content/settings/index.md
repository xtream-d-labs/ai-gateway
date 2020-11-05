+++
title = "Settings"
css = "settings/index.css"
js = "settings/index.js"
+++
<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;margin-bottom: 40px;">
        <form class="form-signin">
          <div class="form-group" style="display: none;">
            <p class="errors"><transition name="fade">
              <span v-if="message != ''">{{ message }}</span>
            </transition></p>
          </div>
          <!-- Docker Registry -->
          <h5 id="docker-registry" class="form-signin-heading"
              style="margin-top: 0;">Docker Registry</h5>
          <hr/>
          <div class="form-group">
            <label for="input-registry" class="row control-label">Registry</label>
            <div class="row">
              <input type="text" id="input-registry" class="form-control" v-model="registry"
                 placeholder="Your private docker registry endpoint (default: DockerHub) " />
            </div>
          </div>
          <div class="form-group">
            <label for="input-hostname" class="row control-label">Hostname</label>
            <div class="row">
              <input type="text" id="input-hostname" class="form-control" v-model="hostname"
                 placeholder="Your private docker registry hostname (default: DockerHub) " />
            </div>
          </div>
          <div class="form-group">
            <label for="input-username" class="row control-label">Username</label>
            <div class="row">
              <input type="text" id="input-username" class="form-control" v-model="username" />
            </div>
          </div>
          <div class="form-group">
            <label for="input-password" class="row control-label">Password</label>
            <div class="row">
              <input type="password" id="input-password" class="form-control" v-model="password" />
            </div>
          </div>
          <!-- NVIDIA GPU Cloud -->
          <h5 id="nvidia-gpu-cloud" class="form-signin-heading">
            <a href="https://ngc.nvidia.com/" target="_blank">NVIDIA GPU Cloud</a>&nbsp;(NGC)
          </h5>
          <hr/>
          <div class="form-group">
            <label for="input-ngc-email" class="row control-label">Email</label>
            <div class="row">
              <input type="text" id="input-ngc-email" class="form-control" v-model="ngcEmail"
                     placeholder="Email address for the web console: foo@bar.com" />
            </div>
          </div>
          <div class="form-group">
            <label for="input-ngc-password" class="row control-label">Password</label>
            <div class="row">
              <input type="password" id="input-ngc-password" class="form-control" v-model="ngcPassword"
                     placeholder="Password for the web console: xxxxx" />
            </div>
          </div>
          <div class="form-group">
            <label for="input-ngc-key" class="row control-label">API Key</label>
            <div class="row">
              <input type="text" id="input-ngc-key" class="form-control" v-model="ngcKey" 
                     placeholder="Generate your API Key at https://ngc.nvidia.com/setup/api-key" />
            </div>
          </div>
          <!-- Local settings -->
          <h5 id="local-settings" class="form-signin-heading">Local settings</h5>
          <hr/>
          <div class="form-group">
            <label for="input-gpus" class="row control-label">GPUs</label>
            <div class="row">
              <input type="text" id="input-gpus" class="form-control" v-model="localGPUs"
                     placeholder="1" readonly="readonly" />
            </div>
          </div>
          <div class="btn btn-lg btn-primary btn-block" v-on:click="submit"
               style="margin: 30px 0 25px 0;">Save</div>
          <div class="clear-both"></div>
        </form>
      </div>
    </div>
  </section>
</main>

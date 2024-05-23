<template>
    <div class="row">
      <div class="col s8 offset-s2" style="margin-top: 50px;">
        <form>
          <div class="input-field col s6">
            <input id="first_name" type="text" class="validate" required>
            <label for="first_name" class="active">First Name</label>
          </div>
          <div class="input-field col s6">
            <input id="last_name" type="text" class="validate" required>
            <label for="last_name" class="active">Last Name</label>
          </div>
        </form>
        <form>
          <div class="input-field col s12">
            <input id="password" type="password" class="validate" required>
            <label for="password" class="active">Password</label>
          </div>
        </form>
        <form>
          <div class="input-field col s12">
            <input id="email" type="email" class="validate" required>
            <label for="email" class="active">Email</label>
          </div>
        </form>
        <div>
          <button class="btn waves-effect grey darken-3" type="submit" name="action">Register
              <i class="material-icons right"></i>
          </button>
      </div>
        <label for="email">Alreay have an account? Sign in</label>
      </div>
    </div>
  </template>
    
  <script>
  </script>
  
  <style>
      .indented{
          margin-left:100px;
      }
  </style>
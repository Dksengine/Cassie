class userAccountTest {
    constructor(userName, userEmail) {
      this.userName = userName;
      this.userEmail = userEmail;
      this.isLoggedIn = false;
    }
  
    logIn() {
      this.isLoggedIn = true;
      console.log(`${this.userName} has logged in.`);
    }
  
    logOut() {
      this.isLoggedIn = false;
      console.log(`${this.userName} has logged out.`);
    }
  
    getUserInfo() {
      return {
        name: this.userName,
        email: this.userEmail,
        status: this.isLoggedIn ? "online" : "offline"
      };
    }
  }

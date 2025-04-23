class user_account_test:
    def __init__(self, user_name, user_email):
        self.user_name = user_name
        self.user_email = user_email
        self.is_logged_in = False

    def log_in(self):
        self.is_logged_in = True
        print(f"{self.user_name} has logged in.")

    def log_out(self):
        self.is_logged_in = False
        print(f"{self.user_name} has logged out.")

    def get_user_info(self):
        return {
            "name": self.user_name,
            "email": self.user_email,
            "status": "online" if self.is_logged_in else "offline"
        }

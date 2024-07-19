package mail

import "fmt"

func GeneratePasswordHTML(to, password string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f4f4f4;
				margin: 0;
				padding: 0;
			}
			.container {
				max-width: 600px;
				margin: 50px auto;
				padding: 20px;
				background-color: #ffffff;
				border-radius: 10px;
				box-shadow: 0 4px 8px rgba(0,0,0,0.1);
				overflow: hidden;
				border: 1px solid #e3e3e3;
			}
			.header {
				background: linear-gradient(to right, #007bff, #00c6ff);
				padding: 20px;
				text-align: center;
				border-radius: 10px;
				border-bottom: 1px solid #e3e3e3;
			}
			.header h1 {
				color: #ffffff;
				margin: 0;
				font-size: 24px;
			}
			.content {
				padding: 20px;
				text-align: center;
			}
			.content p {
				font-size: 16px;
				color: #555555;
				margin: 10px 0;
			}
			.password-container {
				display: flex;
				justify-content: center;
				align-items: center;
				background-color: #f8f8f8;
				padding: 15px;
				border: 1px solid #dddddd;
				border-radius: 5px;
				margin-top: 10px;
			}
			.password {
				font-size: 20px;
				font-weight: bold;
				color: #e67e22;
				margin-right: 10px;
			}
			.copy-button {
				padding: 10px 60px;
				font-size: 18px;
				color: #ffffff;
				background-color:  #8bc34a;
				border: none;
				border-radius: 20px;
				cursor: pointer;
				transition: background-color 0.3s ease;
				box-shadow: 0 4px 8px rgba(0,0,0,0.1);
			}
			.copy-button:hover {
				background-color: #558b2f;
			}
			.footer {
				text-align: center;
				color: #999999;
				font-size: 14px;
				padding: 1px;
				border-top: 1px solid #e3e3e3;
			}
			@media (max-width: 600px) {
				.container {
					margin: 20px;
					padding: 15px;
				}
				.header {
					padding: 15px;
				}
				.content {
					padding: 15px;
				}
				.footer {
					padding: 15px;
				}
			}
		</style>
		<script>
			function copyPassword() {
				var password = document.getElementById('password').textContent;
				var tempInput = document.createElement('input');
				tempInput.value = password;
				document.body.appendChild(tempInput);
				tempInput.select();
				document.execCommand('copy');
				document.body.removeChild(tempInput);
				alert('Password copied to clipboard');
			}
		</script>
	</head>
	<body>
	<div class="container">
		<div class="header">
			<h1>Your New Password</h1>
		</div>
		<div class="content">
			<p>Dear User, your email address is: %s</p>
			<p>Your new password is:</p>
			<br />
			<div class="password-container">
				<p id="password" class="password">%s</p>
			</div>
			<br />
			<button class="copy-button" onclick="copyPassword()">Copy</button>
			<p>Please use this password to log in and remember to change it after your first login.</p>
		</div>
		<div class="footer">
			<p>&copy; 2024 Your Company. All rights reserved.</p>
		</div>
	</div>
</body>
</html>
	`, to, password)
}

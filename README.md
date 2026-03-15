<div align="right">
	<samp>
		<b style="font-size: 14px; color: #c8c2b8;">EN</b> | 
		<a href="README_UA.md" style="text-decoration: none;"><b style="font-size: 13px; color: #8a867f;">UA</b></a> | 
		<a href="README_RU.md" style="text-decoration: none;"><b style="font-size: 13px; color: #8a867f;">RU</b></a>
	</samp>
</div>

<div align="center">
	<h1><samp>[ <span style="color: #5fca08;">UNIT_IDENTIFICATION: LEAVES-MESSENGER</span> ]</samp></h1>
	<p><samp><b style="color: #c8c2b8;">MorangTong Creative Studio</b> // A lightweight, fast, and secure hybrid decentralized messenger.</samp></p>
	<img src="изображение.png" width="100%" style="border: 2px solid #1f2a24; border-radius: 8px; margin-top: 10px; margin-bottom: 20px;" alt="Leaves Messenger Preview">
</div>

<table width="100%" border="0" cellpadding="0" cellspacing="0">
	<tr>
		<td width="50%" valign="top">
			<h3><samp style="color: #c8c2b8;">[ CORE_ENGINE ]</samp></h3>
			<div style="border-left: 2px solid #5fca08; padding-left: 15px; background: #0f1714; padding-top: 10px; padding-bottom: 10px; border-radius: 0 4px 4px 0;">
				<img src="https://img.shields.io/badge/Backend-Go_1.26+-5fca08?style=flat-square&labelColor=0f1714&logo=go&logoColor=c8c2b8">
				<img src="https://img.shields.io/badge/Database-SQLite3-5fca08?style=flat-square&labelColor=0f1714&logo=sqlite&logoColor=c8c2b8"><br>
				<img src="https://img.shields.io/badge/RealTime-Gorilla_WS-5fca08?style=flat-square&labelColor=0f1714&logo=socketdotio&logoColor=c8c2b8">
				<img src="https://img.shields.io/badge/Auth-JWT_&_Bcrypt-5fca08?style=flat-square&labelColor=0f1714&logo=jsonwebtokens&logoColor=c8c2b8">
			</div>
		</td>
		<td width="50%" valign="top">
			<h3><samp style="color: #c8c2b8;">[ CLIENT_INTERFACE ]</samp></h3>
			<div style="border-left: 1px solid #1f2a24; padding-left: 15px; background: #0f1714; padding-top: 10px; padding-bottom: 10px; border-radius: 0 4px 4px 0;">
				<img src="https://img.shields.io/badge/Frontend-Vue_3-5fca08?style=flat-square&labelColor=0f1714&logo=vuedotjs&logoColor=c8c2b8">
				<img src="https://img.shields.io/badge/Language-TypeScript-5fca08?style=flat-square&labelColor=0f1714&logo=typescript&logoColor=c8c2b8"><br>
				<img src="https://img.shields.io/badge/State-Pinia-5fca08?style=flat-square&labelColor=0f1714&logo=vue.js&logoColor=c8c2b8">
				<img src="https://img.shields.io/badge/Runtime-Bun-5fca08?style=flat-square&labelColor=0f1714&logo=bun&logoColor=c8c2b8">
			</div>
		</td>
	</tr>
</table>

<br>

<div align="left">
	<h2><samp style="color: #5fca08;">> system_architecture...</samp></h2>
</div>
<ul style="color: #c8c2b8;">
	<li><samp><b style="color: #5fca08;">[LIGHTNING_FAST]:</b> Zero lag message delivery powered by a custom <b>Go</b> backend utilizing <code>gorilla/websocket</code> for full-duplex real-time communication.</samp></li>
	<li><samp><b style="color: #5fca08;">[SECURE_BY_DESIGN]:</b> End-to-end operational security. Passwords are hashed via <code>x/crypto/bcrypt</code>, sessions are managed by <code>golang-jwt/jwt/v5</code>, and all data is sandboxed in a local <b>SQLite3</b> database.</samp></li>
	<li><samp><b style="color: #5fca08;">[CLEAN_&_MINIMAL]:</b> A modular <b>Vue 3 + TypeScript</b> interface organized into atomic UI kits, seamless composables, and Pinia stores for predictable state management.</samp></li>
	<li><samp><b style="color: #5fca08;">[MODERN_TOOLING]:</b> Next-generation dependency resolution and ultra-fast frontend build speeds powered exclusively by <b>Bun</b>.</samp></li>
</ul>

<br>

<div align="left">
	<h2><samp style="color: #00987a;">> file_structure...</samp></h2>
</div>

<pre style="background-color: #0f1714; padding: 15px; border: 1px solid #1f2a24; border-radius: 6px; color: #c8c2b8; font-family: monospace; font-size: 13px;">
<span style="color: #5fca08;">Leaves/</span>
 ├── <span style="color: #00987a;">backend/</span>                  <span style="color: #6b7a73;"># Go Server Environment</span>
 │    ├── <span style="color: #00987a;">db/</span>                  <span style="color: #6b7a73;"># SQLite database persistence</span>
 │    ├── <span style="color: #00987a;">source/</span>              <span style="color: #6b7a73;"># Core application logic & routers</span>
 │    ├── <span style="color: #c8c2b8;">.env</span> / <span style="color: #c8c2b8;">.env.example</span>  <span style="color: #6b7a73;"># Environment secrets (JWT_KEY, DB_PATH)</span>
 │    ├── <span style="color: #c8c2b8;">go.mod</span> / <span style="color: #c8c2b8;">go.sum</span>      <span style="color: #6b7a73;"># Go dependencies (v1.26.1)</span>
 │    └── <span style="color: #c8c2b8;">main.go</span>              <span style="color: #6b7a73;"># Server entry point</span>
 │
 ├── <span style="color: #00987a;">frontend/</span>                 <span style="color: #6b7a73;"># Vue 3 Client Environment</span>
 │    ├── <span style="color: #00987a;">public/</span>              <span style="color: #6b7a73;"># Static assets</span>
 │    ├── <span style="color: #00987a;">source/</span>
 │    │    ├── <span style="color: #00987a;">api/</span>            <span style="color: #6b7a73;"># Axios/WS clients (auth.ts, chat.ts)</span>
 │    │    ├── <span style="color: #00987a;">components/</span>     <span style="color: #6b7a73;"># Atomic UI modules (ui-kit, AuthForm, ChatInput)</span>
 │    │    ├── <span style="color: #00987a;">composables/</span>    <span style="color: #6b7a73;"># Vue Composition utilities (useTheme.ts)</span>
 │    │    ├── <span style="color: #00987a;">pages/</span>          <span style="color: #6b7a73;"># Main views (AuthPage, HomePage, AdminPage)</span>
 │    │    ├── <span style="color: #00987a;">router/</span>         <span style="color: #6b7a73;"># Vue Router configuration</span>
 │    │    └── <span style="color: #00987a;">stores/</span>         <span style="color: #6b7a73;"># Pinia state (user.ts)</span>
 │    ├── <span style="color: #c8c2b8;">.env</span> / <span style="color: #c8c2b8;">.env.example</span>  <span style="color: #6b7a73;"># API Target URLs (VITE_API_BASE)</span>
 │    └── <span style="color: #c8c2b8;">main.ts</span>              <span style="color: #6b7a73;"># Client mount point</span>
 │
 └── <span style="color: #c8c2b8;">bun.lock</span>                  <span style="color: #6b7a73;"># Universal package lockfile</span>
</pre>

<br>

<div align="left">
	<h2><samp style="color: #8a867f;">> pre_flight_checks...</samp></h2>
</div>
<ul style="color: #c8c2b8;">
	<li><samp><b>Go Compiler:</b> Version <b>1.26.1+</b> is strictly required for backend initialization.</samp></li>
	<li><samp><b>Bun Runtime:</b> Required for installing frontend dependencies, handling linting, and running the dev server.</samp></li>
</ul>

<br>

<div align="left">
	<h2><samp style="color: #5fca08;">> deployment_sequence...</samp></h2>
</div>

<table width="100%" style="text-align: left; border-collapse: collapse; background-color: #0f1714; color: #c8c2b8;">
	<tr style="border-bottom: 2px solid #1f2a24;">
		<th style="padding: 10px; width: 20%;"><samp>Phase</samp></th>
		<th style="padding: 10px;"><samp>Execution Action</samp></th>
	</tr>
	<tr style="border-bottom: 1px solid #1f2a24;">
		<td style="padding: 10px; vertical-align: top;"><samp><b style="color: #5fca08;">01. CLONE_REPO</b></samp></td>
		<td style="padding: 10px;"><samp>Pull the source code to your local machine: <br><code style="background: rgba(95, 202, 8, 0.1); color: #5fca08; padding: 2px 4px; border-radius: 4px;">git clone https://github.com/DrBopoH/leaves.git && cd leaves</code></samp></td>
	</tr>
	<tr style="border-bottom: 1px solid #1f2a24;">
		<td style="padding: 10px; vertical-align: top;"><samp><b style="color: #5fca08;">02. ENV_CONFIG</b></samp></td>
		<td style="padding: 10px;"><samp>Duplicate the example environment files for both subsystems. Adjust values (ports, secrets) in the newly created `.env` files if necessary: <br>
		<code style="background: rgba(95, 202, 8, 0.1); color: #5fca08; padding: 2px 4px; border-radius: 4px;">cp backend/.env.example backend/.env</code><br>
		<code style="background: rgba(95, 202, 8, 0.1); color: #5fca08; padding: 2px 4px; border-radius: 4px;">cp frontend/.env.example frontend/.env</code></samp></td>
	</tr>
	<tr style="border-bottom: 1px solid #1f2a24;">
		<td style="padding: 10px; vertical-align: top;"><samp><b style="color: #5fca08;">03. CORE_INIT</b></samp></td>
		<td style="padding: 10px;"><samp>Navigate to the server directory, download Go modules, and ignite the backend. SQLite database will be generated automatically: <br><code style="background: rgba(95, 202, 8, 0.1); color: #5fca08; padding: 2px 4px; border-radius: 4px;">cd backend && go mod download && go run main.go</code></samp></td>
	</tr>
	<tr style="border-bottom: 1px solid #1f2a24;">
		<td style="padding: 10px; vertical-align: top;"><samp><b style="color: #5fca08;">04. UI_INIT</b></samp></td>
		<td style="padding: 10px;"><samp>In a new terminal split, navigate to the web directory, install packages via Bun, and start the Vite/Vue server: <br><code style="background: rgba(95, 202, 8, 0.1); color: #5fca08; padding: 2px 4px; border-radius: 4px;">cd frontend && bun install && bun run dev</code></samp></td>
	</tr>
	<tr>
		<td style="padding: 10px; vertical-align: top;"><samp><b style="color: #5fca08;">05. ACCESS_PORT</b></samp></td>
		<td style="padding: 10px;"><samp>Open your browser and navigate to the local development port displayed in the Bun terminal (typically <code>http://localhost:5173</code>).</samp></td>
	</tr>
</table>

<br>

<table width="100%" border="0" cellpadding="0" cellspacing="0" style="background-color: #0f1714; border: 1px solid #1f2a24; border-radius: 6px;">
	<tr>
		<td style="padding: 20px;">
			<samp>
				<span style="color: #8a867f;">> Synchronizing local SQLite database...</span><br>
				<span style="color: #5fca08;">[SYSTEM] Leaves WebSocket router active. Listening for connections</span><br>
				<span style="color: #5fca08;">[SYSTEM] Vue 3 Client mounted successfully. Ready for encrypted transmission</span>
				<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=800&size=14&pause=500&color=5fca08&width=20&height=15&vCenter=true&lines=_;%20" alt="cursor" style="vertical-align: middle;">
			</samp>
		</td>
	</tr>
</table>

<br>

<div align="left">
	<h2><samp style="color: #8a867f;">> license_agreement...</samp></h2>
</div>

<pre style="background-color: #0f1714; padding: 15px; border: 1px solid #1f2a24; border-radius: 6px; color: #8a867f; font-family: monospace; font-size: 12px; line-height: 1.4;">
<b style="color: #c8c2b8;">Leaves - A decentralized hybrid messenger.</b>
Copyright (C) 2026 MorangTong Creative Studio

This program is free software: you can redistribute it and/or modify
it under the terms of the <b style="color: #c8c2b8;">GNU Affero General Public License</b> as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see &lt;https://www.gnu.org/licenses/&gt;.
</pre>
package telegram

const (
	// [INFO] start
	msgStart             = `🔒 Welcome to LinkVault, your personal link-saving assistant! 📎`
	msgBotAlreadyStarted = `🤖 LinkVault is already running.`
	// [INFO] greeting
	msgGreeting = msgStart + `


📑➡️📂   With LinkVault, you can easily send us links and we'll save and  organize them.

🔎➡️📧   We'll periodically suggest a random set of saved links or a single one as you prefer, helping you rediscover valuable resources. 

🤔➡️⏰   Set your preferred time interval for link suggestions. 

🤔➡️📝   Edit and manage your saved links effortlessly!

🤔➡️🗑️   Remove links from your list if needed.


⚠️ Beware ⚠️

After you send a link, we put it into your storage, but the message gets deleted immediately.

👉 Use /list to see your saved links.

👉 Use /help to see all available commands with description.`

	// [INFO] help messages
	msgHelpShort = msgStart + `

📑➡️📂  Save your important links just sending (no commands required)

🔎➡️📧  We'll periodically suggest a saved link randomly

🤔➡️⏰  Set time interval for link suggestions

🤔➡️📝  Edit your list of links

🤔➡️🗑️  Remove links from your list

Precautions ⚠️

After you send a link, bot updates your storage, but the message gets deleted immediately.
Use /list to see your saved links.
Use /help see all commands.

👉 Use /help to see all available commands with description.`

	msgHelpWithWarning = `🪄 Currently available commands 🪄

/start - starts the bot


ℹ️ Info ℹ️
/greet  or /g - shows greetings message 

/help /h /commands /cmd - shows this message

/warningdisable, /wd, /w+ - disables warnings about deleting your history

/warningenable, /we, /w- - enables bot to show warnings about deleting your history

⚙️ Settings	⚙️
/list, /l - lists all saved links

/time, /t - sets time interval for link suggestions (default 1 hour)

/remove <link>, /rm <link> - removes a saved link (don't enclose link with <>)

/random, /rnd - instanly suggests a random saved link

/stop - stops the bot, your storage is saved (no sendings from bot until you /start)`

	msgHelpNoWarning = `🪄 Currently available commands 🪄

/start - starts the bot


ℹ️ Info ℹ️

/greet  or /g - shows greetings message 

/help /h /commands /cmd - shows this message

/warningdisable, /wd, /w+ - disables warnings about deleting your history

/warningenable, /we, /w- - enables bot to show warnings about deleting your history

⚙️ Settings ⚙️

/list, /l - lists all saved links

/time, /t - sets time interval for link suggestions (default 1 hour)

/remove <link>, /rm <link> - removes a saved link (don't enclose link with <>)

/random, /rnd - instanly suggests a random saved link

/randomset <number>, /rndset <number> - instanly suggests a random set of links (number - number of links; don't enclose it with <>)

/stop - stops the bot, your storage is saved for 1 month; when timed out, all your saved links will get deleted, so be careful
(no sendings from bot until you /start)



⚠️ Warning ⚠️ Think over before usage! ⚠️

IAmAbsolutelySureIWantToDeletemyHistory - clears your links storage and stops the bot

⬇️ The reason why you can't click on it ⬇️

We care about you, so it's not clickable.
If you're assured of deleting your history, you have to send this command as an ordinary message 
(copy or write by hand without backslash '/' symbol)`

	// [INFO] settings
	msgTimeSetTo = `Time interval is currently set to `

	// [INFO] from readonly
	msgList                          = `Here is your list:`
	msgTimeLeftBeforeStorageDeletion = `Your storage will be deleted in `

	// [ERORR]
	msgUnknownCommand = `Unknown command. Use /help to see all available commands`
	msgListEmpty      = `No saved links. Send a link to save it. Your list is empty. Use /help to see all available commands`
	msgRemoveFailed   = `Link not found. Check if your link is correct.`

	// [INFO] after call
	msgShortcutsEnabled      = `Shortcuts enabled.`
	msgShortcutsDisabled     = `Shortcuts disabled.`
	msgRandom                = `Here is a random link from your list:`
	msgRemovedAll            = `Storage cleared out.`
	msgRemovedSingle         = `Link successfully removed.`
	msgWarningsEnabled       = `Warnings enabled.`
	msgWarningsDisabled      = `Warnings disabled.`
	msgBotStoppedWithTimeout = `Bot stopped. Restart with /start. Your storage is saved for 1 month. When timed out, all your saved links will get deleted. Can't wait for you to come back!`

	// [INFO] after trying to save link
	msgLinkSaved    = `Link successfully saved!`
	msgLinkExists   = `Link already exists in your list.`
	msgNoSavedLinks = `No saved links.`
)

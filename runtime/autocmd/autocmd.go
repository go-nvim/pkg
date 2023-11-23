// Copyright 2023 The Go Nvim Authors
// SPDX-License-Identifier: BSD-3-Clause

// Package autocmd provides the Neovim autocmd.
package autocmd

// CmdlineEnter			After entering the command-line (including
//
//	non-interactive use of ":" in a mapping: use
//	|<Cmd>| instead to avoid this).
//	<afile> expands to the |cmdline-char|.
//	Sets these |v:event| keys:
//	    cmdlevel
//	    cmdtype
//				*CmdlineLeave*
//
// CmdlineLeave			Before leaving the command-line (including
//
//	non-interactive use of ":" in a mapping: use
//	|<Cmd>| instead to avoid this).
//	<afile> expands to the |cmdline-char|.
//	Sets these |v:event| keys:
//	    abort (mutable)
//	    cmdlevel
//	    cmdtype
//	Note: `abort` can only be changed from false
//	to true: cannot execute an already aborted
//	cmdline by changing it to false.
//				*CmdwinEnter*
//
// CmdwinEnter			After entering the command-line window.
//
//	Useful for setting options specifically for
//	this special type of window.
//	<afile> expands to a single character,
//	indicating the type of command-line.
//	|cmdwin-char|
//				*CmdwinLeave*
//
// CmdwinLeave			Before leaving the command-line window.
//
//	Useful to clean up any global setting done
//	with CmdwinEnter.
//	<afile> expands to a single character,
//	indicating the type of command-line.
//	|cmdwin-char|
//				*ColorScheme*
//
// ColorScheme			After loading a color scheme. |:colorscheme|
//
//	The pattern is matched against the
//	colorscheme name. <afile> can be used for the
//	name of the actual file where this option was
//	set, and <amatch> for the new colorscheme
//	name.
//
//				*ColorSchemePre*
//
// ColorSchemePre			Before loading a color scheme. |:colorscheme|
//
//	Useful to setup removing things added by a
//	color scheme, before another one is loaded.
//
// CompleteChanged 					*CompleteChanged*
//
//					After each time the Insert mode completion
//					menu changed.  Not fired on popup menu hide,
//					use |CompleteDonePre| or |CompleteDone| for
//					that.
//
//					Sets these |v:event| keys:
//					    completed_item	See |complete-items|.
//					    height		nr of items visible
//					    width		screen cells
//					    row			top screen row
//					    col			leftmost screen column
//					    size		total nr of items
//					    scrollbar		TRUE if visible
//
//					Non-recursive (event cannot trigger itself).
//					Cannot change the text. |textlock|
//
//	  				The size and position of the popup are also
//	 				available by calling |pum_getpos()|.
//
//								*CompleteDonePre*
//
// CompleteDonePre			After Insert mode completion is done.  Either
//
//	when something was completed or abandoning
//	completion. |ins-completion|
//	|complete_info()| can be used, the info is
//	cleared after triggering CompleteDonePre.
//	The |v:completed_item| variable contains
//	information about the completed item.
//
//				*CompleteDone*
//
// CompleteDone			After Insert mode completion is done.  Either
//
//	when something was completed or abandoning
//	completion. |ins-completion|
//	|complete_info()| cannot be used, the info is
//	cleared before triggering CompleteDone.  Use
//	CompleteDonePre if you need it.
//	|v:completed_item| gives the completed item.
//
//				*CursorHold*
//
// CursorHold			When the user doesn't press a key for the time
//
//	specified with 'updatetime'.  Not re-triggered
//	until the user has pressed a key (i.e. doesn't
//	fire every 'updatetime' ms if you leave Vim to
//	make some coffee. :)  See |CursorHold-example|
//	for previewing tags.
//	This event is only triggered in Normal mode.
//	It is not triggered when waiting for a command
//	argument to be typed, or a movement after an
//	operator.
//	While recording the CursorHold event is not
//	triggered. |q|
//				*<CursorHold>*
//	Internally the autocommand is triggered by the
//	<CursorHold> key. In an expression mapping
//	|getchar()| may see this character.
//
//	Note: Interactive commands cannot be used for
//	this event.  There is no hit-enter prompt,
//	the screen is updated directly (when needed).
//	Note: In the future there will probably be
//	another option to set the time.
//	Hint: to force an update of the status lines
//	use: >
//		:let &ro = &ro
//
// <
//
//	*CursorHoldI*
//
// CursorHoldI			Like CursorHold, but in Insert mode. Not
//
//	triggered when waiting for another key, e.g.
//	after CTRL-V, and not in CTRL-X mode
//	|insert_expand|.
//
//				*CursorMoved*
//
// CursorMoved			After the cursor was moved in Normal or Visual
//
//	mode or to another window.  Also when the text
//	of the cursor line has been changed, e.g. with
//	"x", "rx" or "p".
//	Not triggered when there is typeahead or when
//	an operator is pending.
//	For an example see |match-parens|.
//	Note: Cannot be skipped with |:noautocmd|.
//	Careful: This is triggered very often, don't
//	do anything that the user does not expect or
//	that is slow.
//				*CursorMovedI*
//
// CursorMovedI			After the cursor was moved in Insert mode.
//
//	Not triggered when the popup menu is visible.
//	Otherwise the same as CursorMoved.
//				*DiffUpdated*
//
// DiffUpdated			After diffs have been updated.  Depending on
//
//	what kind of diff is being used (internal or
//	external) this can be triggered on every
//	change or when doing |:diffupdate|.
//				*DirChanged*
//
// DirChanged			After the |current-directory| was changed.
//
//	Sets these |v:event| keys:
//	    cwd:   current working directory
//	    scope: "global", "tab", "window"
//	    changed_window: v:true if we fired the event
//	                    switching window (or tab)
//	Non-recursive (event cannot trigger itself).
//				*FileAppendCmd*
//
// FileAppendCmd			Before appending to a file.  Should do the
//
//	appending to the file.  Use the '[ and ']
//	marks for the range of lines. |Cmd-event|
//				*FileAppendPost*
//
// FileAppendPost			After appending to a file.
//
//	*FileAppendPre*
//
// FileAppendPre			Before appending to a file.  Use the '[ and ']
//
//	marks for the range of lines.
//				*FileChangedRO*
//
// FileChangedRO			Before making the first change to a read-only
//
//	file.  Can be used to checkout the file from
//	a source control system.  Not triggered when
//	the change was caused by an autocommand.
//	Triggered when making the first change in
//	a buffer or the first change after 'readonly'
//	was set, just before the change is applied to
//	the text.
//	WARNING: If the autocommand moves the cursor
//	the effect of the change is undefined.
//				*E788*
//	Cannot switch buffers.  You can reload the
//	buffer but not edit another one.
//				*E881*
//	If the number of lines changes saving for undo
//	may fail and the change will be aborted.
//				*ExitPre*
//
// ExitPre				When using `:quit`, `:wq` in a way it makes
//
//	Vim exit, or using `:qall`, just after
//	|QuitPre|.  Can be used to close any
//	non-essential window.  Exiting may still be
//	cancelled if there is a modified buffer that
//	isn't automatically saved, use |VimLeavePre|
//	for really exiting.
//	See also |QuitPre|, |WinClosed|.
//				*FileChangedShell*
//
// FileChangedShell		When Vim notices that the modification time of
//
//	a file has changed since editing started.
//	Also when the file attributes of the file
//	change or when the size of the file changes.
//	|timestamp|
//	Triggered for each changed file, after:
//	- executing a shell command
//	- |:checktime|
//	- |FocusGained|
//	Not used when 'autoread' is set and the buffer
//	was not changed.  If a FileChangedShell
//	autocommand exists the warning message and
//	prompt is not given.
//	|v:fcs_reason| indicates what happened. Set
//	|v:fcs_choice| to control what happens next.
//	NOTE: Current buffer "%" may be different from
//	the buffer that was changed "<afile>".
//				*E246* *E811*
//	Cannot switch, jump to or delete buffers.
//	Non-recursive (event cannot trigger itself).
//				*FileChangedShellPost*
//
// FileChangedShellPost		After handling a file that was changed outside
//
//	of Vim.  Can be used to update the statusline.
//				*FileReadCmd*
//
// FileReadCmd			Before reading a file with a ":read" command.
//
//	Should do the reading of the file. |Cmd-event|
//				*FileReadPost*
//
// FileReadPost			After reading a file with a ":read" command.
//
//	Note that Vim sets the '[ and '] marks to the
//	first and last line of the read.  This can be
//	used to operate on the lines just read.
//				*FileReadPre*
//
// FileReadPre			Before reading a file with a ":read" command.
//
//	*FileType*
//
// FileType			When the 'filetype' option has been set.  The
//
//	pattern is matched against the filetype.
//	<afile> is the name of the file where this
//	option was set.  <amatch> is the new value of
//	'filetype'.
//	Cannot switch windows or buffers.
//	See |filetypes|.
//				*FileWriteCmd*
//
// FileWriteCmd			Before writing to a file, when not writing the
//
//	whole buffer.  Should do the writing to the
//	file.  Should not change the buffer.  Use the
//	'[ and '] marks for the range of lines.
//	|Cmd-event|
//				*FileWritePost*
//
// FileWritePost			After writing to a file, when not writing the
//
//	whole buffer.
//				*FileWritePre*
//
// FileWritePre			Before writing to a file, when not writing the
//
//	whole buffer.  Use the '[ and '] marks for the
//	range of lines.
//				*FilterReadPost*
//
// FilterReadPost			After reading a file from a filter command.
//
//	Vim checks the pattern against the name of
//	the current buffer as with FilterReadPre.
//	Not triggered when 'shelltemp' is off.
//				*FilterReadPre* *E135*
//
// FilterReadPre			Before reading a file from a filter command.
//
//	Vim checks the pattern against the name of
//	the current buffer, not the name of the
//	temporary file that is the output of the
//	filter command.
//	Not triggered when 'shelltemp' is off.
//				*FilterWritePost*
//
// FilterWritePost			After writing a file for a filter command or
//
//	making a diff with an external diff (see
//	DiffUpdated for internal diff).
//	Vim checks the pattern against the name of
//	the current buffer as with FilterWritePre.
//	Not triggered when 'shelltemp' is off.
//				*FilterWritePre*
//
// FilterWritePre			Before writing a file for a filter command or
//
//	making a diff with an external diff.
//	Vim checks the pattern against the name of
//	the current buffer, not the name of the
//	temporary file that is the output of the
//	filter command.
//	Not triggered when 'shelltemp' is off.
//				*FocusGained*
//
// FocusGained			Nvim got focus.
//
//	*FocusLost*
//
// FocusLost			Nvim lost focus.  Also (potentially) when
//
//	a GUI dialog pops up.
//				*FuncUndefined*
//
// FuncUndefined			When a user function is used but it isn't
//
//	defined.  Useful for defining a function only
//	when it's used.  The pattern is matched
//	against the function name.  Both <amatch> and
//	<afile> are set to the name of the function.
//	NOTE: When writing Vim scripts a better
//	alternative is to use an autoloaded function.
//	See |autoload-functions|.
//				*UIEnter*
//
// UIEnter				After a UI connects via |nvim_ui_attach()|,
//
//	after VimEnter.  Can be used for GUI-specific
//	configuration.
//	Sets these |v:event| keys:
//	    chan
//				*UILeave*
//
// UILeave				After a UI disconnects from Nvim.
//
//	Sets these |v:event| keys:
//	    chan
//				*InsertChange*
//
// InsertChange			When typing <Insert> while in Insert or
//
//	Replace mode.  The |v:insertmode| variable
//	indicates the new mode.
//	Be careful not to move the cursor or do
//	anything else that the user does not expect.
//				*InsertCharPre*
//
// InsertCharPre			When a character is typed in Insert mode,
//
//	before inserting the char.
//	The |v:char| variable indicates the char typed
//	and can be changed during the event to insert
//	a different character.  When |v:char| is set
//	to more than one character this text is
//	inserted literally.
//
//	Cannot change the text. |textlock|
//	Not triggered when 'paste' is set.
//				*TextYankPost*
//
// TextYankPost			Just after a |yank| or |deleting| command, but not
//
//	if the black hole register |quote_| is used nor
//	for |setreg()|. Pattern must be *.
//	Sets these |v:event| keys:
//	    inclusive
//	    operator
//	    regcontents
//	    regname
//	    regtype
//	    visual
//	The `inclusive` flag combined with the |'[|
//	and |']| marks can be used to calculate the
//	precise region of the operation.
//
//	Non-recursive (event cannot trigger itself).
//	Cannot change the text. |textlock|
//				*InsertEnter*
//
// InsertEnter			Just before starting Insert mode.  Also for
//
//	Replace mode and Virtual Replace mode.  The
//	|v:insertmode| variable indicates the mode.
//	Be careful not to do anything else that the
//	user does not expect.
//	The cursor is restored afterwards.  If you do
//	not want that set |v:char| to a non-empty
//	string.
//				*InsertLeavePre*
//
// InsertLeavePre			Just before leaving Insert mode.  Also when
//
//	using CTRL-O |i_CTRL-O|.  Be caseful not to
//	change mode or use `:normal`, it will likely
//	cause trouble.
//				*InsertLeave*
//
// InsertLeave			Just after leaving Insert mode.  Also when
//
//	using CTRL-O |i_CTRL-O|.  But not for |i_CTRL-C|.
//				*MenuPopup*
//
// MenuPopup			Just before showing the popup menu (under the
//
//	right mouse button).  Useful for adjusting the
//	menu for what is under the cursor or mouse
//	pointer.
//	The pattern is matched against a single
//	character representing the mode:
//		n	Normal
//		v	Visual
//		o	Operator-pending
//		i	Insert
//		c	Command line
//				*OptionSet*
//
// OptionSet			After setting an option (except during
//
//	|startup|).  The |autocmd-pattern| is matched
//	against the long option name.  |<amatch>|
//	indicates what option has been set.
//
//	|v:option_type| indicates whether it's global
//	or local scoped.
//	|v:option_command| indicates what type of
//	set/let command was used (follow the tag to
//	see the table).
//	|v:option_new| indicates the newly set value.
//	|v:option_oldlocal| has the old local value.
//	|v:option_oldglobal| has the old global value.
//	|v:option_old| indicates the old option value.
//
//	|v:option_oldlocal| is only set when |:set|
//	or |:setlocal| or a |modeline| was used to set
//	the option. Similarly |v:option_oldglobal| is
//	only set when |:set| or |:setglobal| was used.
//
//	Note that when setting a |global-local| string
//	option with |:set|, then |v:option_old| is the
//	old global value. However, for all other kinds
//	of options (local string options, global-local
//	number options, ...) it is the old local
//	value.
//
//	OptionSet is not triggered on startup and for
//	the 'key' option for obvious reasons.
//
//	Usage example: Check for the existence of the
//	directory in the 'backupdir' and 'undodir'
//	options, create the directory if it doesn't
//	exist yet.
//
//	Note: Do not reset the same option during this
//	autocommand, that may break plugins. You can
//	always use |:noautocmd| to prevent triggering
//	OptionSet.
//
//	Non-recursive: |:set| in the autocommand does
//	not trigger OptionSet again.
//
//				*QuickFixCmdPre*
//
// QuickFixCmdPre			Before a quickfix command is run (|:make|,
//
//	|:lmake|, |:grep|, |:lgrep|, |:grepadd|,
//	|:lgrepadd|, |:vimgrep|, |:lvimgrep|,
//	|:vimgrepadd|, |:lvimgrepadd|, |:cscope|,
//	|:cfile|, |:cgetfile|, |:caddfile|, |:lfile|,
//	|:lgetfile|, |:laddfile|, |:helpgrep|,
//	|:lhelpgrep|, |:cexpr|, |:cgetexpr|,
//	|:caddexpr|, |:cbuffer|, |:cgetbuffer|,
//	|:caddbuffer|).
//	The pattern is matched against the command
//	being run.  When |:grep| is used but 'grepprg'
//	is set to "internal" it still matches "grep".
//	This command cannot be used to set the
//	'makeprg' and 'grepprg' variables.
//	If this command causes an error, the quickfix
//	command is not executed.
//				*QuickFixCmdPost*
//
// QuickFixCmdPost			Like QuickFixCmdPre, but after a quickfix
//
//	command is run, before jumping to the first
//	location. For |:cfile| and |:lfile| commands
//	it is run after error file is read and before
//	moving to the first error.
//	See |QuickFixCmdPost-example|.
//				*QuitPre*
//
// QuitPre				When using `:quit`, `:wq` or `:qall`, before
//
//	deciding whether it closes the current window
//	or quits Vim.  Can be used to close any
//	non-essential window if the current window is
//	the last ordinary window.
//	See also |ExitPre|, ||WinClosed|.
//				*RemoteReply*
//
// RemoteReply			When a reply from a Vim that functions as
//
//	server was received |server2client()|.  The
//	pattern is matched against the {serverid}.
//	<amatch> is equal to the {serverid} from which
//	the reply was sent, and <afile> is the actual
//	reply string.
//	Note that even if an autocommand is defined,
//	the reply should be read with |remote_read()|
//	to consume it.
//				*SessionLoadPost*
//
// SessionLoadPost			After loading the session file created using
//
//	the |:mksession| command.
//				*ShellCmdPost*
//
// ShellCmdPost			After executing a shell command with |:!cmd|,
//
//	|:make| and |:grep|.  Can be used to check for
//	any changed files.
//	For non-blocking shell commands, see
//	|job-control|.
//				*Signal*
//
// Signal				After Nvim receives a signal. The pattern is
//
//	matched against the signal name. Only
//	"SIGUSR1" is supported.  Example: >
//	    autocmd Signal SIGUSR1 call some#func()
//
// <							*ShellFilterPost*
// ShellFilterPost			After executing a shell command with
//
//	":{range}!cmd", ":w !cmd" or ":r !cmd".
//	Can be used to check for any changed files.
//				*SourcePre*
//
// SourcePre			Before sourcing a Vim script. |:source|
//
//	<afile> is the name of the file being sourced.
//				*SourcePost*
//
// SourcePost			After sourcing a Vim script. |:source|
//
//	<afile> is the name of the file being sourced.
//	Not triggered when sourcing was interrupted.
//	Also triggered after a SourceCmd autocommand
//	was triggered.
//				*SourceCmd*
//
// SourceCmd			When sourcing a Vim script. |:source|
//
//	<afile> is the name of the file being sourced.
//	The autocommand must source this file.
//	|Cmd-event|
//				*SpellFileMissing*
//
// SpellFileMissing		When trying to load a spell checking file and
//
//	it can't be found.  The pattern is matched
//	against the language.  <amatch> is the
//	language, 'encoding' also matters.  See
//	|spell-SpellFileMissing|.
//				*StdinReadPost*
//
// StdinReadPost			During startup, after reading from stdin into
//
//	the buffer, before executing modelines. |--|
//				*StdinReadPre*
//
// StdinReadPre			During startup, before reading from stdin into
//
//	the buffer. |--|
//				*SwapExists*
//
// SwapExists			Detected an existing swap file when starting
//
//	to edit a file.  Only when it is possible to
//	select a way to handle the situation, when Vim
//	would ask the user what to do.
//	The |v:swapname| variable holds the name of
//	the swap file found, <afile> the file being
//	edited.  |v:swapcommand| may contain a command
//	to be executed in the opened file.
//	The commands should set the |v:swapchoice|
//	variable to a string with one character to
//	tell Vim what should be done next:
//		'o'	open read-only
//		'e'	edit the file anyway
//		'r'	recover
//		'd'	delete the swap file
//		'q'	quit, don't edit the file
//		'a'	abort, like hitting CTRL-C
//	When set to an empty string the user will be
//	asked, as if there was no SwapExists autocmd.
//				*E812*
//	Cannot change to another buffer, change
//	the buffer name or change directory.
//				*Syntax*
//
// Syntax				When the 'syntax' option has been set.  The
//
//	pattern is matched against the syntax name.
//	<afile> expands to the name of the file where
//	this option was set. <amatch> expands to the
//	new value of 'syntax'.
//	See |:syn-on|.
//				*TabEnter*
//
// TabEnter			Just after entering a tab page. |tab-page|
//
//	After WinEnter.
//	Before BufEnter.
//				*TabLeave*
//
// TabLeave			Just before leaving a tab page. |tab-page|
//
//	After WinLeave.
//				*TabNew*
//
// TabNew				When creating a new tab page. |tab-page|
//
//	After WinEnter.
//	Before TabEnter.
//				*TabNewEntered*
//
// TabNewEntered			After entering a new tab page. |tab-page|
//
//	After BufEnter.
//				*TabClosed*
//
// TabClosed			After closing a tab page. <afile> expands to
//
//	the tab page number.
//				*TermOpen*
//
// TermOpen			When a |terminal| job is starting.  Can be
//
//	used to configure the terminal buffer.
//				*TermEnter*
//
// TermEnter			After entering |Terminal-mode|.
//
//	After TermOpen.
//				*TermLeave*
//
// TermLeave			After leaving |Terminal-mode|.
//
//	After TermClose.
//				*TermClose*
//
// TermClose			When a |terminal| job ends.
//
//	*TermResponse*
//
// TermResponse			After the response to t_RV is received from
//
//	the terminal.  The value of |v:termresponse|
//	can be used to do things depending on the
//	terminal version.  May be triggered halfway
//	through another event (file I/O, a shell
//	command, or anything else that takes time).
//				*TextChanged*
//
// TextChanged			After a change was made to the text in the
//
//	current buffer in Normal mode.  That is after
//	|b:changedtick| has changed (also when that
//	happened before the TextChanged autocommand
//	was defined).
//	Not triggered when there is typeahead or when
//	an operator is pending.
//	Note: Cannot be skipped with `:noautocmd`.
//	Careful: This is triggered very often, don't
//	do anything that the user does not expect or
//	that is slow.
//				*TextChangedI*
//
// TextChangedI			After a change was made to the text in the
//
//	current buffer in Insert mode.
//	Not triggered when the popup menu is visible.
//	Otherwise the same as TextChanged.
//				*TextChangedP*
//
// TextChangedP			After a change was made to the text in the
//
//	current buffer in Insert mode, only when the
//	popup menu is visible.  Otherwise the same as
//	TextChanged.
//				*User*
//
// User				Not executed automatically.  Use |:doautocmd|
//
//	to trigger this, typically for "custom events"
//	in a plugin.  Example: >
//	    :autocmd User MyPlugin echom 'got MyPlugin event'
//	    :doautocmd User MyPlugin
//
// <							*UserGettingBored*
// UserGettingBored		When the user presses the same key 42 times.
//
//	Just kidding! :-)
//				*VimEnter*
//
// VimEnter			After doing all the startup stuff, including
//
//	loading vimrc files, executing the "-c cmd"
//	arguments, creating all windows and loading
//	the buffers in them.
//	Just before this event is triggered the
//	|v:vim_did_enter| variable is set, so that you
//	can do: >
//	   if v:vim_did_enter
//	     call s:init()
//	   else
//	     au VimEnter * call s:init()
//	   endif
//
// <							*VimLeave*
// VimLeave			Before exiting Vim, just after writing the
//
//	.shada file.  Executed only once, like
//	VimLeavePre.
//	Use |v:dying| to detect an abnormal exit.
//	Use |v:exiting| to get the exit code.
//	Not triggered if |v:dying| is 2 or more.
//				*VimLeavePre*
//
// VimLeavePre			Before exiting Vim, just before writing the
//
//	.shada file.  This is executed only once,
//	if there is a match with the name of what
//	happens to be the current buffer when exiting.
//	Mostly useful with a "*" pattern. >
//	   :autocmd VimLeavePre * call CleanupStuff()
//
// <				Use |v:dying| to detect an abnormal exit.
//
//	Use |v:exiting| to get the exit code.
//	Not triggered if |v:dying| is 2 or more.
//				*VimResized*
//
// VimResized			After the Vim window was resized, thus 'lines'
//
//	and/or 'columns' changed.  Not when starting
//	up though.
//				*VimResume*
//
// VimResume			After Nvim resumes from |suspend| state.
//
//	*VimSuspend*
//
// VimSuspend			Before Nvim enters |suspend| state.
//
//	*WinClosed*
//
// WinClosed			After closing a window. <afile> expands to the
//
//	|window-ID|.
//	After WinLeave.
//	Non-recursive (event cannot trigger itself).
//	See also |ExitPre|, |QuitPre|.
//				*WinEnter*
//
// WinEnter			After entering another window.  Not done for
//
//	the first window, when Vim has just started.
//	Useful for setting the window height.
//	If the window is for another buffer, Vim
//	executes the BufEnter autocommands after the
//	WinEnter autocommands.
//	Note: For split and tabpage commands the
//	WinEnter event is triggered after the split
//	or tab command but before the file is loaded.
//
//				*WinLeave*
//
// WinLeave			Before leaving a window.  If the window to be
//
//	entered next is for a different buffer, Vim
//	executes the BufLeave autocommands before the
//	WinLeave autocommands (but not for ":new").
//	Not used for ":qa" or ":q" when exiting Vim.
//	After WinClosed.
//				*WinNew*
//
// WinNew				When a new window was created.  Not done for
//
//	the first window, when Vim has just started.
//	Before WinEnter.
//				*WinScrolled*
//
// WinScrolled			After scrolling the viewport of the current
//
//	window.

// List of autocmd events.
const (
	// BufAdd Just after creating a new buffer which is added to the buffer list, or adding a buffer to the buffer list, a buffer in the buffer list was renamed.
	//
	// Before BufEnter.
	BufAdd = "BufAdd"

	// BufDelete before deleting a buffer from the buffer list.
	BufDelete = "BufDelete"

	// BufEnter after entering a buffer.
	//
	// Useful for setting options for a file type.
	// Also executed when starting to edit a buffer.
	//
	// after BufAdd.
	// after BufReadPost.
	BufEnter = "BufEnter"

	// BufFilePost after changing the name of the current buffer with the ":file" or ":saveas" command.
	BufFilePost = "BufFilePost"

	// BufFilePre before changing the name of the current buffer with the ":file" or ":saveas" command.
	BufFilePre = "BufFilePre"

	// BufHidden before a buffer becomes hidden: when there are no longer windows that show the buffer, but the buffer is not unloaded or deleted.
	BufHidden = "BufHidden"

	// BufLeave before leaving to another buffer.
	//
	// Also when leaving or closing the current window and the new current window is not for the same buffer.
	BufLeave = "BufLeave"

	// BufModifiedSet After the "modified" value of a buffer has been changed.
	//
	// This autocmd Neovim specific.
	BufModifiedSet = "BufModifiedSet"

	// BufNew Just after creating a new buffer.
	// Also used just after a buffer has been renamed.
	//
	// When the buffer is added to the buffer list BufAdd will be triggered too.
	BufNew = "BufNew"

	// BufNewFile When starting to edit a file that doesn't exist.
	// Can be used to read in a skeleton file.
	BufNewFile = "BufNewFile"

	// BufReadPost starting to edit a new buffer, after reading the file.
	BufReadPost = "BufReadPost"

	// BufRead starting to edit a new buffer, after reading the file.
	//
	// Alias of BufReadPost.
	BufRead = BufReadPost

	// BufReadCmd Before starting to edit a new buffer.
	// Should read the file into the buffer. Cmd-event.
	BufReadCmd = "BufReadCmd"

	// BufReadPre when starting to edit a new buffer, after reading the file into the buffer, before processing modelines.
	//
	// See BufWinEnter to do something after processing modelines.
	//
	// Also triggered:
	//  when writing an unnamed buffer such that the buffer gets a name
	//  after successfully recovering a file
	//  for the "filetypedetect" group when executing ":filetype detect"
	//
	// Not triggered:
	//  for ":r file"
	//  if the file doesn't exist
	BufReadPre = "BufReadPre"

	// BufUnload before unloading a buffer, when the text in the buffer is going to be freed.
	//
	// After BufWritePost.
	// Before BufDelete.
	BufUnload = "BufUnload"

	// BufWinEnter After a buffer is displayed in a window.
	//
	// This may be when the buffer is loaded (after processing modelines) or when a hidden buffer is displayed (and is no longer hidden).
	BufWinEnter = "BufWinEnter"

	// BufWinLeave before a buffer is removed from a window.
	//
	// Not when it's still visible in another window.
	// Also triggered when exiting.
	//
	// Before BufUnload, BufHidden.
	BufWinLeave = "BufWinLeave"

	// BufWipeout before completely deleting a buffer.
	//
	// The BufUnload and BufDelete events may be called first (if the buffer was loaded and was in the buffer list).
	//
	// Also used just before a buffer is renamed (also when it's not in the buffer list).
	BufWipeout = "BufWipeout"

	// BufWritePre before writing the whole buffer to a file.
	BufWritePre = "BufWritePre"

	// BufWrite before writing the whole buffer to a file.
	//
	// Alias of BufWritePre.
	BufWrite = BufWritePre

	// BufWriteCmd Before writing the whole buffer to a file.
	BufWriteCmd = "BufWriteCmd"

	// BufWritePost after writing the whole buffer to a file (should undo the commands for BufWritePre).
	BufWritePost = "BufWritePost"

	// ChanInfo state of channel changed, for instance the client of a RPC channel described itself.
	// Sets these |v:event| keys: info.
	// See |nvim_get_chan_info()| for the format of the info Dictionary.
	ChanInfo = "ChanInfo"

	// ChanOpen just after a channel was opened.
	// sets these |v:event| keys: info.
	//
	// See |nvim_get_chan_info()| for the format of the info Dictionary.
	ChanOpen = "ChanOpen"

	// CmdUndefined when a user command is used but it isn't defined.
	// Useful for defining a command only when it's used.
	//
	// The pattern is matched against the command name.
	// Both <amatch> and <afile> expand to the command name.
	CmdUndefined = "CmdUndefined"

	// CmdlineChanged after a change was made to the text inside command line.
	// Be careful not to mess up the command line, it may cause Vim to lock up. <afile> expands to the cmdline-char.
	CmdlineChanged = "CmdlineChanged"
)

// List of Reading autocmd name.
const (
	// FileReadPre before reading a file with a ":read" command.
	FileReadPre = "FileReadPre"

	// FileReadPost after reading a file with a ":read" command.
	FileReadPost = "FileReadPost"

	// FileReadCmd before reading a file with a ":read" command. See also `:help Cmd-event`.
	FileReadCmd = "FileReadCmd"

	// FilterReadPre before reading a file from a filter command.
	FilterReadPre = "FilterReadPre"

	// FilterReadPost after reading a file from a filter command.
	FilterReadPost = "FilterReadPost"

	// StdinReadPre before reading from stdin into the buffer.
	StdinReadPre = "StdinReadPre"

	// StdinReadPost After reading from the stdin into the buffer.
	StdinReadPost = "StdinReadPost"
)

// List of Writing autocmd name.
const (
	// FileWritePre starting to write part of a buffer to a file.
	FileWritePre = "FileWritePre"

	// FileWritePost after writing part of a buffer to a file.
	FileWritePost = "FileWritePost"

	// FileWriteCmd before writing part of a buffer to a file. See also `:help Cmd-event`.
	FileWriteCmd = "FileWriteCmd"

	// FileAppendPre starting to append to a file.
	FileAppendPre = "FileAppendPre"

	// FileAppendPost after appending to a file.
	FileAppendPost = "FileAppendPost"

	// FileAppendCmd before appending to a file. See also `:help Cmd-event`.
	FileAppendCmd = "FileAppendCmd"

	// FilterWritePre starting to write a file for a filter command or diff.
	FilterWritePre = "FilterWritePre"

	// FilterWritePost after writing a file for a filter command or diff.
	FilterWritePost = "FilterWritePost"
)

// List of Buffers autocmd name.
const (
	// BufCreate just after adding a buffer to the buffer list.
	//
	// Alias of BufAdd.
	BufCreate = BufAdd

	// SwapExists detected an existing swap file.
	SwapExists = "SwapExists"
)

// List of Options autocmd name.
const (
	// FileType when the 'filetype' option has been set.
	FileType = "FileType"

	// Syntax when the 'syntax' option has been set.
	Syntax = "Syntax"

	// OptionSet after setting any option Startup and exit.
	OptionSet = "OptionSet"

	// VimEnter after doing all the startup stuff.
	VimEnter = "VimEnter"

	// GUIEnter after starting the GUI successfully.
	GUIEnter = "GUIEnter"

	// GUIFailed after starting the GUI failed.
	GUIFailed = "GUIFailed"

	// TermResponse after the terminal response to t_RV is received.
	TermResponse = "TermResponse"

	// QuitPre when using `:quit`, before deciding whether to exit.
	QuitPre = "QuitPre"

	// ExitPre when using a command that may make Vim exit.
	ExitPre = "ExitPre"

	// VimLeavePre before exiting Nvim, before writing the shada file.
	VimLeavePre = "VimLeavePre"

	// VimLeave before exiting Nvim, after writing the shada file.
	VimLeave = "VimLeave"

	// VimResume after Nvim is resumed.
	VimResume = "VimResume"

	// VimSuspend before Nvim is suspended.
	VimSuspend = "VimSuspend"
)

// List of Various autocmd name.
const (
	// DiffUpdated after diffs have been updated.
	DiffUpdated = "DiffUpdated"

	// FileChangedShell Vim notices that a file changed since editing started.
	FileChangedShell = "FileChangedShell"

	// FileChangedShellPost after handling a file changed since editing started.
	FileChangedShellPost = "FileChangedShellPost"

	// FileChangedRO before making the first change to a read-only file.
	FileChangedRO = "FileChangedRO"

	// ShellCmdPost after executing a shell command.
	ShellCmdPost = "ShellCmdPost"

	// ShellFilterPostafter filtering with a shell command.
	ShellFilterPostafter = "ShellFilterPostafter"

	// FuncUndefined a user function is used but it isn't defined.
	FuncUndefined = "FuncUndefined"

	// SpellFileMissing a spell file is used but it can't be found.
	SpellFileMissing = "SpellFileMissing"

	// SourcePre before sourcing a Vim script.
	SourcePre = "SourcePre"

	// SourceCmd before sourcing a Vim script |Cmd-event|.
	SourceCmd = "SourceCmd"

	// VimResized after the Vim window size changed.
	VimResized = "VimResized"

	// FocusGained Nvim got focus.
	FocusGained = "FocusGained"

	// FocusLost Nvim lost focus.
	FocusLost = "FocusLost"

	// CursorHold the user doesn't press a key for a while.
	CursorHold = "CursorHold"

	// CursorHoldI the user doesn't press a key for a while in Insert mode.
	CursorHoldI = "CursorHoldI"

	// CursorMoved the cursor was moved in Normal mode.
	CursorMoved = "CursorMoved"

	// CursorMovedI the cursor was moved in Insert mode WinNew after creating a new window.
	CursorMovedI = "CursorMovedI"

	// WinEnter after entering another window.
	WinEnter = "WinEnter"

	// WinScrolled after scrolling the viewport of the current window.
	//
	// This autocmd Neovim specific.
	WinScrolled = "WinScrolled"

	// WinLeavet before leaving a window.
	WinLeavet = "WinLeavet"

	// WinClosed after closing a window. <afile> expands to the window-ID. after WinLeave.
	//
	// This autocmd Neovim specific.
	WinClosed = "WinClosed"

	// TabNew when creating a new tab page.
	//
	// This autocmd Neovim specific.
	TabNew = "TabNew"

	// TabNewEntered after entering a new tab page.
	//
	// This autocmd Neovim specific.
	TabNewEntered = "TabNewEntered"

	// TabEntert after entering another tab page.
	TabEntert = "TabEntert"

	// TabLeavet before leaving a tab page.
	TabLeavet = "TabLeavet"

	// TabClosed after closing a tab page.
	//
	// This autocmd Neovim specific.
	TabClosed = "TabClosed"

	// CmdlineEnter after entering cmdline mode.
	CmdlineEnter = "CmdlineEnter"

	// CmdlineLeave before leaving cmdline mode.
	CmdlineLeave = "CmdlineLeave"

	// CmdwinEnter after entering the command-line window.
	CmdwinEnter = "CmdwinEnter"

	// CmdwinLeave before leaving the command-line window.
	CmdwinLeave = "CmdwinLeave"

	// InsertEnter starting Insert mode.
	InsertEnter = "InsertEnter"

	// InsertChange when typing <Insert> while in Insert or Replace mode.
	InsertChange = "InsertChange"

	// InsertLeave when leaving Insert mode.
	InsertLeave = "InsertLeave"

	// InsertCharPre when a character was typed in Insert mode, before inserting it.
	InsertCharPre = "InsertCharPre"

	// TextYankPost when some text is yanked or deleted.
	TextYankPost = "TextYankPost"

	// TextChanged after a change was made to the text in Normal mode.
	TextChanged = "TextChanged"

	// TextChangedI after a change was made to the text in Insert mode when popup menu is not visible.
	TextChangedI = "TextChangedI"

	// TextChangedP after a change was made to the text in Insert mode when popup menu visible.
	TextChangedP = "TextChangedP"

	// ColorSchemePre before loading a color scheme.
	ColorSchemePre = "ColorSchemePre"

	// ColorScheme after loading a color scheme.
	ColorScheme = "ColorScheme"

	// RemoteReply a reply from a server Vim was received.
	RemoteReply = "RemoteReply"

	// QuickFixCmdPre before a quickfix command is run.
	QuickFixCmdPre = "QuickFixCmdPre"

	// QuickFixCmdPost after a quickfix command is run.
	QuickFixCmdPost = "QuickFixCmdPost"

	// SessionLoadPost after loading a session file.
	SessionLoadPost = "SessionLoadPost"

	// MenuPopup just before showing the popup menu.
	MenuPopup = "MenuPopup"

	// CompleteChanged after popup menu changed, not fired on popup menu hide.
	CompleteChanged = "CompleteChanged"

	// CompleteDone after Insert mode completion is done.
	CompleteDone = "CompleteDone"

	// DirChanged after the `current-directory` was changed.
	//
	// This autocmd Neovim specific.
	DirChanged = "DirChanged"

	// Signal after Nvim receives a signal.
	//
	// This autocmd Neovim specific.
	Signal = "Signal"

	// User to be used in combination with ":doautocmd".
	User = "User"
)

// List of terminal autocmd name.
const (
	// TermOpen when a terminal job starts.
	//
	// This autocmd Neovim specific.
	TermOpen = "TermOpen"

	// TermEnter after entering Terminal mode. after TermOpen.
	//
	// This autocmd Neovim specific.
	TermEnter = "TermEnter"

	// TermLeave after leaving Terminal mode.
	TermLeave = "TermLeave"

	// TermClose when a terminal job ends.
	//
	// This autocmd Neovim specific.
	TermClose = "TermClose"
)

// List of UD autocmd name.
const (
	// UIEnter after a UI connects via nvim_ui_attach(), after VimEnter. Can be used for GUI-specific configuration.
	//
	// This autocmd Neovim specific.
	UIEnter = "UIEnter"

	// UILeave after a UI disconnects from Nvim.
	//
	// This autocmd Neovim specific.
	UILeave = "UILeave"
)

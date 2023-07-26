package main

import (
	"fmt"
	"image/color"
	"path/filepath"
	"log"
	"time"
	"errors"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func validateFolderName(s string) error {
	if len(s) == 0 {
		return errors.New("name not found")
	}
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case `<`, `>`, `:`, `"`, `/`, `\`, `|`, `?`, `*`:
			return errors.New(`unsupported character`)
		}
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func move_files(path string, 
		types *widget.SelectEntry,
	 	name *widget.Entry,
	  	existing_dirs []os.DirEntry,
	   	existing_files []os.DirEntry,
	    age int, 
		chek_age bool) (bool, string) {

	// -----------Validate folder name-----------
	if name.Validate() != nil {
		return false, "bad folder name"
	}
	//----------- Check if folder already exists-----------
	new_folder := true
	for _, dir := range existing_dirs {
		if dir.Name() == name.Text {
			new_folder = false
		}
	}

	// ----------- Set some helping vars and create new_dir_name-----------
	files_not_yet_found := true
	new_dir_name := path + string(os.PathSeparator) + name.Text
	touch := true

	// ----------- loop through files-----------
	for _, fl := range existing_files {
		//----------- if need to check age that find if to touch a file-----------
		if chek_age{
			file_info, _ := os.Stat(path + string(os.PathSeparator) + fl.Name())
			last_mod := file_info.ModTime()
			border_date := time.Now().AddDate(0, 0, -age)
			touch = border_date.Before(last_mod)
		}

		// ----------- if file is right format, and touch -> mkdir and move file
		if strings.ToLower(filepath.Ext(fl.Name())) == types.Text && touch{
			if files_not_yet_found && new_folder {
				os.Mkdir(new_dir_name, os.ModePerm)
			}
			files_not_yet_found = false
			old_path := path + string(os.PathSeparator) + fl.Name()
			new_path := new_dir_name + string(os.PathSeparator) + fl.Name()
			os.Rename(old_path, new_path)
		}
	}
	// -----------if files were not moved -> false -----------
	if files_not_yet_found {
		return false, "files not found"
	}
	return true, "success"
}


func main() {

	// -----------Building Main App-----------
	log.SetOutput(nil)
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("SeaSorter")
	message_color := color.NRGBA{R: 255, G: 100, B: 100, A: 255}
	myWindow.Resize(fyne.Size{
		Width:  460,
		Height: 500,
	})
	myIcon, _ := fyne.LoadResourceFromURLString(`https://iili.io/HQDor7e.png`)
	welcome_text := canvas.NewText("Welcome to SeaSorter - easy way to sort your files by format", color.NRGBA{R: 153, G: 204, B: 255, A: 255})
	welcome_text.TextStyle = fyne.TextStyle{Italic: true}
	welcome := container.NewCenter(welcome_text)
	myWindow.SetIcon(myIcon)

	// ------------Report------------
	log_of_result := container.NewVBox()
	report := container.New(layout.NewMaxLayout(), log_of_result)
	report.Hide()

	// ------------Folder Path------------
	input := widget.NewEntry()
	input.SetPlaceHolder(`Enter the path to the desired folder (ex. c:\Users\Me\Desktop)`)
	input.OnChanged = func(s string) {
		report.Hide()
	}

	// ------------Example words------------
	instruction_1 := canvas.NewText("Create partitions like this", color.NRGBA{R: 153, G: 204, B: 255, A: 255})
	instruction_1.TextStyle = fyne.TextStyle{Italic: true}
	example := container.NewCenter(instruction_1)

	// Example - Icon of the folder and wiriting - Name
	file_icon := widget.NewIcon(theme.FileIcon())
	folder_icon := widget.NewIcon(theme.FolderIcon())
	real_name_writing := container.New(layout.NewCenterLayout(),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(60, 60)), folder_icon), canvas.NewText("Photos", color.Black))
	real_folder_icon := container.New(layout.NewCenterLayout(),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)), file_icon), canvas.NewText(".jpg", color.Black))

	fields := container.New(layout.NewGridLayout(2), real_folder_icon, real_name_writing)

	// -----------Separators-----------
	separator := widget.NewSeparator()
	separator_2 := widget.NewSeparator()

	//-----------------First entry from user type_of_the_file & name_of_the_folder----------------------
	moreEvilNinjas := []string{".aac", ".aiff", ".avi", ".bmp", ".c", ".cpp", ".csv", ".dat", ".dmg", ".dll", ".dng", ".doc", ".docx", ".eps", ".exe", ".flac", ".gif", ".h", ".html", ".ics", ".iso", ".java", ".jpeg", ".jpg", ".json", ".key", ".log", ".m4a", ".mp3", ".mp4", ".mpg", ".odt", ".otf", ".pdf", ".png", ".ppt", ".pptx", ".psd", ".py", ".rar", ".raw", ".rtf", ".svg", ".tar", ".tex", ".tga", ".tif", ".tiff", ".ts", ".txt", ".wav", ".webm", ".wmv", ".xls", ".xlsx", ".xml", ".yaml", ".zip"}
	folder_type := widget.NewSelectEntry(moreEvilNinjas)
	folder_type.SetPlaceHolder("Enter file type (ex .jpg)")

	folder_name := widget.NewEntry()
	folder_name.SetPlaceHolder("Enter folder name")

	folder_type.OnChanged = func(s string) {
		report.Hide()
	}
	folder_name.OnChanged = func(s string) {
		report.Hide()
	}

	// add validator to folder name
	folder_name.Validator = validateFolderName

	// create grid
	one_folder := container.New(layout.NewGridLayout(2), folder_type, folder_name)
	grid := container.New(layout.NewGridLayout(1), one_folder)




	//-----------Delete and Add buttons-----------
	// del button
	delete_button := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		if len(grid.Objects) > 1 {
			grid.Objects = grid.Objects[:len(grid.Objects)-1]
		}
	})
	// create button
	create_button := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		f_type := widget.NewSelectEntry(moreEvilNinjas)
		d_name := widget.NewEntry()
		f_type.OnChanged = func(s string) {
			report.Hide()
		}
		d_name.OnChanged = func(s string) {
			report.Hide()
		}
		grid.Add(container.New(layout.NewGridLayout(2), f_type, d_name))
		d_name.Validator = validateFolderName
		grid.Refresh()
	})
	control_buttons := container.New(layout.NewGridLayout(2), delete_button, create_button)

	instruction_2_t := canvas.NewText("Select filters if desired", color.NRGBA{R: 153, G: 204, B: 255, A: 255})
	instruction_2_t.TextStyle = fyne.TextStyle{Italic: true}
	instruction_2 := container.NewCenter(instruction_2_t)

	//-----Check auto------
	check_tic := widget.NewCheck("Automatic sorting", func(sth bool) {
		if sth {
			grid.Hide()
			control_buttons.Hide()
			example.Hide()
			fields.Hide()
		} else {
			grid.Show()
			control_buttons.Show()
			example.Show()
			fields.Show()
		}
	})

	//----Slider---
	days_count := widget.NewLabel("0 days and younger")
	slider := widget.NewSlider(float64(1), float64(30))
	days_restriction := container.New(layout.NewFormLayout(), days_count, slider)
	days_restriction.Hide()


	//-----Check date------
	check_date := widget.NewCheck("Filter by date", func(oth bool) {
		if oth {
			days_restriction.Show()
		} else {
			days_restriction.Hide()
		}
	})

	//--------------Control slider while changing----------------
	slider.OnChanged = func(value float64) {
		days_count.SetText(fmt.Sprintf("%.0f days and younger", value))
	}


	//------Progress bar-------
	progress := widget.NewProgressBarInfinite()
	progress.Hide()


	//-----Checkers----
	checkers := container.New(layout.NewFormLayout(), check_tic, check_date)


	//-----------Execute Button Action-----------
	execute := widget.NewButton("Execute", func() {
		// set progres bar at start
		progress.Show()
		// clear log
		log_of_result.Objects = log_of_result.Objects[:0]
		// check if path exists
		if exists(input.Text) {
			// find existing files and dirs
			var existing_dirs []os.DirEntry
			var existing_files []os.DirEntry
			entries, err := os.ReadDir(input.Text)
			if err != nil {
				log.Fatal(err)
			}
			for _, e := range entries {
				if e.IsDir() {
					existing_dirs = append(existing_dirs, e)
				} else {
					existing_files = append(existing_files, e)
				}
			}

			// Line of automatic sorting
			if check_tic.Checked{
				// create partitions
				dct_of_types_and_dirs := make(map[string]string)
				for _, i := range existing_files{
					dct_of_types_and_dirs[filepath.Ext(i.Name())] = strings.ToUpper(filepath.Ext(i.Name())[1:])
				}

				// sort files into partitions
				for f_t_1, d_n_1 := range dct_of_types_and_dirs {

					// creating mock-up objects
					file_type_1 := widget.NewSelectEntry(moreEvilNinjas)
					file_type_1.SetText(f_t_1)
					dir_name_1 := widget.NewEntry()
					dir_name_1.SetText(d_n_1)

					// execute main functions
					success, message := move_files(input.Text,file_type_1, dir_name_1,existing_dirs, existing_files,int(slider.Value), check_date.Checked)

					// decide on color of one log
					if success {
						message_color = color.NRGBA{R: 100, G: 255, B: 100, A: 255}
					} else {
						message_color = color.NRGBA{R: 255, G: 100, B: 100, A: 255}
					}

					log_of_result.Add(canvas.NewText(file_type_1.Text + " -> " + dir_name_1.Text + ":  " + message, message_color))
				}

			// Line of manual sorting
			} else{
			for i := 0; i < len(grid.Objects); i++ {

				// get objects from containers
				my_container := grid.Objects[i].(*fyne.Container)
				file_type_1 := my_container.Objects[0].(*widget.SelectEntry)
				dir_name_1 := my_container.Objects[1].(*widget.Entry)

				// execute main function
				done, message := move_files(
					input.Text,
					file_type_1, dir_name_1,
					existing_dirs, existing_files,
					int(slider.Value), check_date.Checked)
				
				// decide on color of one log
				if done {
					message_color = color.NRGBA{R: 100, G: 255, B: 100, A: 255}
				} else {
					message_color = color.NRGBA{R: 255, G: 100, B: 100, A: 255}
				}

				// add massage into log
				log_of_result.Add(canvas.NewText(file_type_1.Text + " -> " + dir_name_1.Text + ":  " + message, message_color))
			}}
		} else {
			log_of_result.Add(canvas.NewText("path does not exist", message_color))
		}
		report.Show()
		progress.Hide()
	})

	// Add all elements to one container
	content := container.NewVBox(welcome, progress, input,
		separator, example, fields, grid, control_buttons, separator_2,
		instruction_2, checkers, days_restriction, execute, report)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

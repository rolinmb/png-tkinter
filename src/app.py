import subprocess
import tkinter as tk

select_options = ["Generator 1", "Generator 2", "Generator 3", "Generator 4"]

class App:
    def __init__(self, root):
        self.root = root
        self.root.title("Tkinter App")
        self.menu_frame = tk.Frame(self.root)
        self.menu_frame.pack()
        
        self.home_button = tk.Button(self.menu_frame, text="Home", command=self.show_home)
        self.generate_button = tk.Button(self.menu_frame, text="Generate", command=self.show_generate)
        
        self.home_button.pack(side="left", padx=10, pady=10)
        self.generate_button.pack(side="left", padx=10, pady=10)
        
        self.home_frame = tk.Frame(self.root)
        self.generate_frame = tk.Frame(self.root)
        
        self.home_label = tk.Label(self.home_frame, text="Welcome to the Home Page")
        self.generate_label = tk.Label(self.generate_frame, text="Welcome to the PNG Generator")
        
        self.gen_cfactor_label = tk.Label(self.generate_frame, text="Color Factor: 100")
        self.gen_cfactor_label.pack(pady=10)
        
        self.gen_cfactor_var = tk.IntVar()
        self.gen_cfactor_var.set(100)
        
        self.gen_cfactor_slider = tk.Scale(self.generate_frame, from_=1, to=500, resolution=1, variable=self.gen_cfactor_var, command=self.update_gen_cfactor)
        self.gen_cfactor_slider.pack()
        
        self.gen_geo_label = tk.Label(self.generate_frame, text="Geometry Size: 50")
        self.gen_geo_label.pack(pady=10)
        
        self.gen_geo_var = tk.IntVar()
        self.gen_geo_var.set(50)
        
        self.gen_geo_slider = tk.Scale(self.generate_frame, from_=1, to=100, resolution=1, variable=self.gen_geo_var, command=self.update_gen_geo)
        self.gen_geo_slider.pack()
        
        self.gen_scale_label = tk.Label(self.generate_frame, text="Scale: 0.005")
        self.gen_scale_label.pack(pady=10)
        
        self.gen_scale_var = tk.DoubleVar()
        self.gen_scale_var.set(0.005)
        
        self.gen_scale_slider = tk.Scale(self.generate_frame, from_=0, to=1, resolution=0.001, variable=self.gen_scale_var, command=self.update_gen_scale)
        self.gen_scale_slider.pack()
        
        self.gen_chaos_label = tk.Label(self.generate_frame, text="Chaos: 0.025")
        self.gen_chaos_label.pack(pady=10)
        
        self.gen_chaos_var = tk.DoubleVar()
        self.gen_chaos_var.set(0.25)
        
        self.gen_chaos_slider = tk.Scale(self.generate_frame, from_=0, to=0.5, resolution=0.001, variable=self.gen_chaos_var, command=self.update_gen_chaos)
        self.gen_chaos_slider.pack()
        
        self.gen_selected = tk.StringVar()
        self.gen_selected.set('Generator 1')
        
        self.gen_selector = tk.OptionMenu(self.generate_frame, self.gen_selected, *select_options)
        self.gen_selector.pack()
        
        self.gen_fname_input = tk.Text(self.generate_frame, height=2.5, width=20)
        self.gen_fname_input.pack()
        
        self.run_button = tk.Button(self.generate_frame, text="Generate new .png", command=self.generate)
        self.run_button.pack()
        
        self.show_home()

    def run(self):
        self.root.mainloop()

    def show_home(self):
        self.hide_all_frames()
        self.home_frame.pack()
        self.home_label.pack()
    
    def show_generate(self):
        self.hide_all_frames()
        self.generate_frame.pack()
        self.generate_label.pack()
    
    def hide_all_frames(self):
        self.home_frame.pack_forget()
        self.generate_frame.pack_forget()
  
    def update_gen_cfactor(self, _=None):
        val = self.gen_cfactor_var.get()
        self.gen_cfactor_label.config(text=f"Color Factor: {val}")
  
    def update_gen_geo(self, _=None):
        val = self.gen_geo_var.get()
        self.gen_geo_label.config(text=f"Geometry Size: {val}")
    
    def update_gen_scale(self, _=None):
        val = self.gen_scale_var.get()
        self.gen_scale_label.config(text=f"Scale: {val}")
    
    def update_gen_chaos(self, _=None):
        val = self.gen_chaos_var.get()
        self.gen_chaos_label.config(text=f"Chaos: {val}")
    
    def generate(self):
        clrFactor = self.gen_cfactor_var.get()
        geoVal = self.gen_geo_var.get()
        scaleVal = self.gen_scale_var.get()
        chaosVal = self.gen_chaos_var.get()
        fout = self.gen_fname_input.get(1.0, "end-1c")
        opt = self.gen_selected.get()
        subprocess.run(["go", "run", "main.go", fout, str(clrFactor), str(geoVal), str(scaleVal), str(chaosVal), opt])
  
if __name__ == "__main__":
  pass
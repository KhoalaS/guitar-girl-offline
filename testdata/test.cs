[Serializable]
public class initDataInfo : TBase, TAbstractBase
{
	// Fields
	private Int16 _idx; // 0x10
	private String _game_url; // 0x18
	private String _cdn_url; // 0x20
	public Isset __isset; // 0x28

	// Properties
	
	// Methods
	// RVA: 0x1e87f70 VA: 0x753a393f70
	public Int16 get_Idx() { }
	// RVA: 0x1e87f78 VA: 0x753a393f78
	public Void set_Idx(Int16 value) { }
	// RVA: 0x1e87f88 VA: 0x753a393f88
	public String get_Game_url() { }
	// RVA: 0x1e87f90 VA: 0x753a393f90
	public Void set_Game_url(String value) { }
	// RVA: 0x1e87fa0 VA: 0x753a393fa0
	public String get_Cdn_url() { }
	// RVA: 0x1e87fa8 VA: 0x753a393fa8
	public Void set_Cdn_url(String value) { }
	// RVA: 0x1e87fb8 VA: 0x753a393fb8
	public Void .ctor() { }
	// RVA: 0x1e87fc0 VA: 0x753a393fc0
	public Void Read(TProtocol iprot) { }
	// RVA: 0x1e88218 VA: 0x753a394218
	public Void Write(TProtocol oprot) { }
	// RVA: 0x1e88540 VA: 0x753a394540
	public override String ToString() { }
}

[Serializable]
public class initRetDataInfo : TBase, TAbstractBase
{
	// Fields
	private Int16 _idx; // 0x10
	private String _game_url; // 0x18
	private String _cdn_url; // 0x20
	public Isset __isset; // 0x28

	// Properties
	public Int16 Idx { get; set; }
	public String Game_url { get; set; }
	public String Cdn_url { get; set; }

	// Methods
	// RVA: 0x1e87f70 VA: 0x753a393f70
	public Int16 get_Idx() { }
	// RVA: 0x1e87f78 VA: 0x753a393f78
	public Void set_Idx(Int16 value) { }
	// RVA: 0x1e87f88 VA: 0x753a393f88
	public String get_Game_url() { }
	// RVA: 0x1e87f90 VA: 0x753a393f90
	public Void set_Game_url(String value) { }
	// RVA: 0x1e87fa0 VA: 0x753a393fa0
	public String get_Cdn_url() { }
	// RVA: 0x1e87fa8 VA: 0x753a393fa8
	public Void set_Cdn_url(String value) { }
	// RVA: 0x1e87fb8 VA: 0x753a393fb8
	public Void .ctor() { }
	// RVA: 0x1e87fc0 VA: 0x753a393fc0
	public Void Read(TProtocol iprot) { }
	// RVA: 0x1e88218 VA: 0x753a394218
	public Void Write(TProtocol oprot) { }
	// RVA: 0x1e88540 VA: 0x753a394540
	public override String ToString() { }
}

public class initRetDataInfo : TBase, TAbstractBase
{
	[SOME CODE]
}